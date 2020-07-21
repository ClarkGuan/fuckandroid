#!/usr/bin/env bash

rice clean
rice embed-go
go install ./cmd/fuckandroid
