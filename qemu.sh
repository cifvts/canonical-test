#!/bin/bash

set -e

REPOSITORY="https://cloud-images.ubuntu.com/kinetic/current/"
TARGZ_IMAGE="kinetic-server-cloudimg-amd64.tar.gz"
BASE_IMAGE="kinetic-server-cloudimg-amd64.img"
QEMU_IMAGE="ubuntu-kinetic-qemu.img"
CLOUD_INIT_YAML="ubuntu-cloud-init.yaml"
CLOUD_INIT_IMAGE="ubuntu-cloud-init.img"

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit
fi

if [[ ! -f ${TARGZ_IMAGE} ]]; then
	wget ${REPOSITORY}${TARGZ_IMAGE}
fi

if [[ ! -f ${BASE_IMAGE} ]]; then
	tar -xf ${TARGZ_IMAGE} ${BASE_IMAGE}
fi

if [[ ! -f ${QEMU_IMAGE} ]]; then
	qemu-img create -f qcow2 -b ${BASE_IMAGE} ${QEMU_IMAGE}
fi

cloud-localds ${CLOUD_INIT_IMAGE} ${CLOUD_INIT_YAML}

qemu-system-x86_64 \
-machine accel=kvm \
-kernel /boot/vmlinuz \
-append "root=/dev/sda single console=ttyS0 systemd.unit=graphical.target" \
-hda ${QEMU_IMAGE} \
-hdb ${CLOUD_INIT_IMAGE} \
-m 2048 \
--nographic \
-netdev user,id=net0,hostfwd=tcp::2222-:22 -device virtio-net-pci,netdev=net0 \
