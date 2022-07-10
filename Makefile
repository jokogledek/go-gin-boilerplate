#!/bin/bash
export PKGS=$(shell go list ./...)
export GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct
export GO111MODULE=on