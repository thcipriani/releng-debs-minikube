## minikube ISO image

This includes the configuration for an alternative bootable ISO image meant to be used in conjection with minikube.

It includes:
- systemd as the init system
- rkt
- docker
- CRI-O

## Hacking

### Requirements

* Linux
```
sudo apt-get install build-essential gnupg2 p7zip-full git wget cpio python \
	unzip bc gcc-multilib automake libtool locales
```

Either import your private key or generate a sign-only key using `gpg2 --gen-key`.
Also be sure to have an UTF-8 locale set up in order to build the ISO.

### Build instructions

```
$ git clone https://github.com/kubernetes/minikube
$ cd minikube
$ make buildroot-image
$ make out/minikube.iso
```

The build will occurs inside a docker container, if you want to do this
baremetal, replace `make out/minikube.iso` with `IN_DOCKER=1 make out/minikube.iso`.
The bootable ISO image will be available in `out/minikube.iso`.

### Testing local minikube-iso changes

```
$ ./out/minikube start \
    --container-runtime=rkt \
    --network-plugin=cni \
    --iso-url=file://$GOPATH/src/k8s.io/minikube/out/minikube.iso
```

### Buildroot configuration

To change the buildroot configuration, execute:

```
$ cd out/buildroot
$ make menuconfig
$ make
```

To save any buildroot configuration changes made with `make menuconfig`, execute:

```
$ cd out/buildroot
$ make savedefconfig
```

The changes will be reflected in the `minikube-iso/configs/minikube_defconfig` file.

```
$ git status
## master
 M deploy/iso/minikube-iso/configs/minikube_defconfig
```

### Saving buildroot/kernel configuration changes


To make any kernel configuration changes and save them, execute:

```
$ make linux-menuconfig
```

This will open the kernel configuration menu, and then save your changes to our
iso directory after they've been selected.
