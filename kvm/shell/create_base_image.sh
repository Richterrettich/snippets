#!/bin/sh

./create_disk.sh

virt-install \
    --name $1 \
    --memory 4000 \
    --vcpus=2,maxvcpus=4 \
    --disk ../out/shell/test.img,size=4,format=qcow2,device=disk,bus=virtio \
    --network bridge=bridge0 \
    --cpu host \
    --virt-type kvm \
    --os-type=linux \
    --os-variant generic  \
    --graphics vnc,password=foobar,port=5910,keymap=de \
    --cd $2
