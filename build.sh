#!/bin/sh
set -e

gcc -c cmain.c
go build -v -buildmode=plugin -o plugin.so plugin.go
go build -v -o gomain gomain.go
