#!/bin/bash

set -euo pipefail

go clean
go generate
go build
./modprox-registry ../../hack/configs/registry-local.postgres.json
