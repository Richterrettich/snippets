#!/bin/sh


virt-install \
  --name demo2 \
  --memory 4000 \
  --vcpus=2,maxvcpus=4 \
  --disk ./test2.img,size=4,format=qcow2,device=disk,bus=virtio \
  --network bridge=bridge0 \
  --cpu host \
  --virt-type kvm \
  --os-type=linux \
  --os-variant centos7.0  \
  --graphics vnc,password=foobar,port=5910,keymap=de 
#  --cd $HOME/Downloads/ubuntu-16.04.1-server-amd64.iso \ 
#  --location 'http://mirror.i3d.net/pub/centos/7/os/x86_64/' \
#  --console pty,target_type=serial \
#  --extra-args 'console=ttyS0,115200n8 serial'
