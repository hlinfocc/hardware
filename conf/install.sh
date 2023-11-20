#!/bin/bash

# Check if user is root
if [ $(id -u) != "0" ]; then
    echo "Error: You must be root to run this script, please use root to install"
    exit 1
fi

if [ -f "./hlinfo-hardware" ];then
    \cp ./hlinfo-hardware /usr/local/bin/hlinfo-hardware
    \cp ./hlinfo-hardware.service /etc/systemd/system/
    systemctl enable hlinfo-hardware
    systemctl start hlinfo-hardware
fi
