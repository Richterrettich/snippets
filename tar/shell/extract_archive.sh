#!/bin/sh

OUTPUT_DIR=out

mkdir -p $OUTPUT_DIR && tar -Jxf test.tar.xz -C $OUTPUT_DIR --strip-components 1 
