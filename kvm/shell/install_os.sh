#!/bin/sh
qemu-system-x86_64 -cdrom ~/Downloads/archlinux-2016.12.01-dual.iso -boot order=d -drive format=qcow2,file=../out/shell/test.img -m 4G -enable-kvm -net nic -net bridge,br=bridge0
