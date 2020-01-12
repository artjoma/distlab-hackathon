#!/usr/bin/env bash

echo "start wasm gen"
eosio-cpp -O3 -abigen src/claims.cpp -o build/claims.wasm
echo "end"

