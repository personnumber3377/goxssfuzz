#!/bin/sh

export PATH=/home/oof/.asdf/installs/golang/1.23.0/packages/bin:$PATH

go-fuzz-build -libfuzzer

clang -fsanitize=fuzzer ./reflect-fuzz.a -o fuzzer

