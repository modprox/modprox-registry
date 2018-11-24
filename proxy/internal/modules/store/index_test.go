package store

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/modprox/mp/pkg/coordinates"
	"github.com/modprox/mp/pkg/repository"

	"github.com/stretchr/testify/require"
)

func setupIndex(t *testing.T) (string, Index) {
	tmpDir, err := ioutil.TempDir("", "index-")
	require.NoError(t, err)

	options := IndexOptions{
		Directory: tmpDir,
	}

	index, err := NewIndex(options)
	require.NoError(t, err)

	return tmpDir, index
}

func cleanupIndex(t *testing.T, tmpDir string) {
	err := os.RemoveAll(tmpDir)
	require.NoError(t, err)
}

func newMod(source, version string) coordinates.Module {
	return coordinates.Module{
		Source:  source,
		Version: version,
	}
}

func Test_Index_empty(t *testing.T) {
	tmpDir, index := setupIndex(t)
	defer cleanupIndex(t, tmpDir)

	versions, err := index.Versions("github.com/pkg/errors")
	require.NoError(t, err)
	require.Equal(t, 0, len(versions))

	_, err = index.Info(newMod(
		"github.com/pkg/errors",
		"v0.8.0",
	))
	require.Error(t, err)

	_, err = index.Mod(newMod(
		"github.com/pkg/errors",
		"v0.8.0",
	))
	require.Error(t, err)

	exists, _, err := index.Contains(newMod(
		"github.com/pkg/errors",
		"v0.8.0",
	))
	require.NoError(t, err)
	require.False(t, exists)
}

func Test_Index_Put_1(t *testing.T) {
	tmpDir, index := setupIndex(t)
	defer cleanupIndex(t, tmpDir)

	err := index.Put(ModuleAddition{
		Mod: coordinates.Module{
			Source:  "github.com/pkg/errors",
			Version: "v0.8.0",
		},
		ModFile:  "module github.com/pkg/errors",
		UniqueID: 1,
	})
	require.NoError(t, err)

	versions, err := index.Versions("github.com/pkg/errors")
	require.NoError(t, err)
	require.Equal(t, 1, len(versions))

	// not the module added
	versions, err = index.Versions("gitlab.com/some/other")
	require.NoError(t, err)
	require.Equal(t, 0, len(versions))

	info, err := index.Info(newMod(
		"github.com/pkg/errors",
		"v0.8.0",
	))
	require.NoError(t, err)
	require.Equal(t, repository.RevInfo{
		Version: "v0.8.0",
	}, info)

	// not the module added
	_, err = index.Info(newMod(
		"github.com/pkg/errors",
		"v6.6.6",
	))
	require.Error(t, err)

	modFile, err := index.Mod(newMod(
		"github.com/pkg/errors",
		"v0.8.0",
	))
	require.NoError(t, err)
	require.Equal(t, "module github.com/pkg/errors", modFile)

	// not the module added
	_, err = index.Mod(newMod(
		"github.com/pkg/errors",
		"v6.6.6",
	))
	require.Error(t, err)
}

// todo: test were we put several in, and test version sorting

func insert(t *testing.T, index Index, pkg string, id int) {
	err := index.Put(ModuleAddition{
		Mod: coordinates.Module{
			Source:  pkg,
			Version: fmt.Sprintf("v0.0.%d", id),
		},
		ModFile:  fmt.Sprintf("module %s", pkg),
		UniqueID: int64(id),
	})
	require.NoError(t, err)
}

func Test_IDs_empty(t *testing.T) {
	tmpDir, index := setupIndex(t)
	defer cleanupIndex(t, tmpDir)

	ids, err := index.IDs()
	require.NoError(t, err)
	require.Equal(t, Ranges(nil), ids)
}

func Test_IDs(t *testing.T) {
	tmpDir, index := setupIndex(t)
	defer cleanupIndex(t, tmpDir)

	insert(t, index, "github.com/pkg/errors", 1)
	insert(t, index, "github.com/pkg/errors", 2)
	insert(t, index, "github.com/pkg/errors", 3)
	insert(t, index, "github.com/pkg/errors", 4)
	insert(t, index, "github.com/pkg/errors", 5)

	insert(t, index, "github.com/pkg/toolkit", 10)
	insert(t, index, "github.com/pkg/toolkit", 11)
	insert(t, index, "github.com/pkg/errors", 12)

	insert(t, index, "github.com/pkg/errors", 20)

	ids, err := index.IDs()
	require.NoError(t, err)
	require.Equal(t, Ranges{
		{1, 5}, {10, 12}, {20, 20},
	}, ids)
}

func Test_ranges(t *testing.T) {
	try := func(input []int64, exp Ranges) {
		output := ranges(input)
		require.Equal(t, exp, output)
	}

	try(
		[]int64{},
		Ranges(nil),
	)

	try(
		[]int64{5},
		Ranges{{5, 5}},
	)

	try(
		[]int64{7, 8},
		Ranges{{7, 8}},
	)

	try(
		[]int64{2, 3, 4, 7, 8, 10, 13},
		Ranges{{2, 4}, {7, 8}, {10, 10}, {13, 13}},
	)

	try(
		[]int64{0, 4, 5, 6, 7, 8, 23, 25, 26},
		Ranges{{0, 0}, {4, 8}, {23, 23}, {25, 26}},
	)
}

func Test_first(t *testing.T) {
	try := func(input []int64, expRange Range, expLen int) {
		incRange, lenRange := first(input)
		require.Equal(t, expRange, incRange)
		require.Equal(t, expLen, lenRange)
	}

	try(
		[]int64{},
		Range{0, 0}, 0,
	)

	try(
		[]int64{5},
		Range{5, 5}, 1,
	)

	try(
		[]int64{7, 8},
		Range{7, 8}, 2,
	)

	try(
		[]int64{3, 6},
		Range{3, 3}, 1,
	)

	try(
		[]int64{3, 4, 5, 6, 8, 9, 10},
		Range{3, 6}, 4,
	)

	try(
		[]int64{4, 7, 8, 9},
		Range{4, 4}, 1,
	)
}

func Test_Summary_empty(t *testing.T) {
	tmpDir, index := setupIndex(t)
	defer cleanupIndex(t, tmpDir)

	mods, versions, err := index.Summary()
	require.NoError(t, err)
	require.Equal(t, 0, mods)
	require.Equal(t, 0, versions)
}

func Test_Summary(t *testing.T) {
	tmpDir, index := setupIndex(t)
	defer cleanupIndex(t, tmpDir)

	checkSummary(t, index, 0, 0)

	insert(t, index, "github.com/pkg/errors", 1)
	insert(t, index, "github.com/pkg/errors", 2)
	insert(t, index, "github.com/pkg/errors", 3)
	insert(t, index, "github.com/pkg/errors", 4)
	insert(t, index, "github.com/pkg/errors", 5)

	insert(t, index, "github.com/pkg/toolkit", 10)
	insert(t, index, "github.com/pkg/toolkit", 11)
	insert(t, index, "github.com/pkg/errors", 12)

	insert(t, index, "github.com/pkg/errors", 20)

	checkSummary(t, index, 2, 9)
}

func checkSummary(t *testing.T, index Index, expMods, expVers int) {
	mods, versions, err := index.Summary()
	require.NoError(t, err)
	require.Equal(t, expMods, mods)
	require.Equal(t, expVers, versions)
}

func Test_Remove(t *testing.T) {
	tmpDir, index := setupIndex(t)
	defer cleanupIndex(t, tmpDir)

	checkSummary(t, index, 0, 0)

	insert(t, index, "github.com/pkg/errors", 1)
	insert(t, index, "github.com/pkg/errors", 2)
	insert(t, index, "github.com/pkg/errors", 3)
	insert(t, index, "github.com/pkg/errors", 4)
	insert(t, index, "github.com/pkg/errors", 5)
	insert(t, index, "github.com/pkg/toolkit", 10)
	insert(t, index, "github.com/pkg/toolkit", 11)
	insert(t, index, "github.com/pkg/errors", 12)
	insert(t, index, "github.com/pkg/errors", 20)

	checkSummary(t, index, 2, 9)

	m1 := coordinates.Module{
		Source:  "github.com/pkg/toolkit",
		Version: "v0.0.10",
	}

	exists, id, err := index.Contains(m1)
	require.NoError(t, err)
	require.True(t, exists)
	require.Equal(t, int64(10), id)

	err = index.Remove(m1)
	require.NoError(t, err)

	exists, _, err = index.Contains(m1)
	require.NoError(t, err)
	require.False(t, exists)

	checkSummary(t, index, 2, 8)
}
