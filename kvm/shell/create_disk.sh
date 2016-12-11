#!/bin/sh
mkdir -p ../out/shell
qemu-img create -f qcow2 ./test.img 10G
