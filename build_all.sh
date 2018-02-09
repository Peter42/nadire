#!/bin/bash
set -e

rm -rf build
mkdir build

# all supported versions https://golang.org/doc/install/source#environment as of 09.02.2018
#targets="android_arm darwin_386 darwin_amd64 darwin_arm darwin_arm64 dragonfly_amd64 freebsd_386 freebsd_amd64 freebsd_arm linux_386 linux_amd64 linux_arm linux_arm64 linux_ppc64 linux_ppc64le linux_mips linux_mipsle linux_mips64 linux_mips64le netbsd_386 netbsd_amd64 netbsd_arm openbsd_386 openbsd_amd64 openbsd_arm plan9_386 plan9_amd64 solaris_amd64 windows_386 windows_amd64"

# feel free to create a pull request for additional targets
targets="linux_386 linux_amd64 linux_arm linux_arm64 windows_386 windows_amd64 darwin_386 darwin_amd64"

function get_extension() {
	case $target_os in
	"windows")
		target_extension=".exe"
		;;
	*)
		target_extension=""
		;;
	esac
}



for target in $targets; do
	
	target_arch=$(echo $target | sed 's:^[a-z0-9]*_::' )
	target_os=$(echo $target | sed 's:_[a-z0-9]*::' )
	get_extension
	
	export GOOS=$target_os
	export GOARCH=$target_arch
	target="build/nadire-$target_os-$target_arch$target_extension"
	echo "Building $target"
	go build -o $target
done