#!/usr/bin/env bash
set -ex
PROJECT_DIR=$(realpath $(dirname $0)/..)

migrate -database "postgres://postgres:postgrespw@127.0.0.1:5432/helloer?sslmode=disable" -path ${PROJECT_DIR}/migrations/ -verbose "$@"