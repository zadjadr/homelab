#!/bin/bash
#
# Init. script for cloud VM
# Cleans iptables and makes the changes persistant
sudo iptables -F
sudo iptables -nvL
sudo netfilter-persistent save
sudo systemctl status  netfilter-persistent
sudo systemctl enable netfilter-persistent
