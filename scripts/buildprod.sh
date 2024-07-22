#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 cd src && go build -o pokedex
