#!/bin/sh
mkdir -p ../out/shell
qemu-img create -f qcow2 ../out/shell/test.img 10G
