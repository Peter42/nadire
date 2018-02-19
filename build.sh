if [[ $# -eq 2 ]]; then
	target_os=$1
	target_arch=$2
	export GOOS=$target_os
	export GOARCH=$target_arch
	
	target_extension=`go env GOEXE`
	target="build/nadire-$target_os-$target_arch$target_extension"
else
	target_extension=`go env GOEXE`
	target="nadire$target_extension"
fi

echo "Building $target"
go build -o $target src/main.go