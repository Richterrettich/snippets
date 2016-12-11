#!/bin/sh


virt-install \
  --name demo \
  --memory 4000 \
  --vcpus=2,maxvcpus=4 \
  --disk ./test.img,size=4,format=qcow2,device=disk,bus=virtio \
  --network bridge=bridge0 \
  --cpu host \
  --virt-type kvm \
  --os-type=linux \
  --os-variant centos7.0  \
  --cd $HOME/Downloads/archlinux-2016.12.01-dual.iso
#  --location 'http://mirror.i3d.net/pub/centos/7/os/x86_64/' \
  --console pty,target_type=serial \
  --graphics vnc,password=foobar,port=5910,keymap=de \
  --extra-args 'console=ttyS0,115200n8 serial'
