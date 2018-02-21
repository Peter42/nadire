#!/bin/bash
find src/ -name *.go | xargs -L 1 gofmt -w -l -s
