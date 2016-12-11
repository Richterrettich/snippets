#!/bin/sh

rm ../out/shell/test.img
virsh destroy demo
virsh undefine demo
