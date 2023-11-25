# compile for version
make
if [ $? -ne 0 ]; then
    echo "make error"
    exit 1
fi

hardware_version=`./bin/hlinfo-hardware -v`
echo "build version: ${hardware_version}"

if [ -d "./dist" ];then
    rm -rf ./dist
fi

mkdir -p ./dist/packages

# cross_compiles
make -f ./Makefile.cross-compiles


os_all='linux freebsd windows'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle riscv64'

cd ./dist

for os in $os_all; do
    for arch in $arch_all; do
        echo "package: OS:${os},arch: ${arch}"
        hardware_dir_name="hlinfo-hardware_${hardware_version}_${os}_${arch}"
        hardware_path="./packages/hlinfo-hardware_${hardware_version}_${os}_${arch}"

        if [ "x${os}" = x"windows" ];then
            if [ ! -f "./hlinfo-hardware_${os}_${arch}.exe" ];then
                continue
            fi
            mkdir -p ${hardware_path}
            mv ./hlinfo-hardware_${os}_${arch}.exe ${hardware_path}/hlinfo-hardware.exe
        else
             if [ ! -f "./hlinfo-hardware_${os}_${arch}" ];then
                continue
            fi
            mkdir -p ${hardware_path}
            mv ./hlinfo-hardware_${os}_${arch} ${hardware_path}/hlinfo-hardware
        fi

        cp ../LICENSE ${hardware_path}
        if [ "x${os}" = x"linux" ]; then
            \cp ../conf/linux/* ${hardware_path}/
        elif [ "x${os}" = x"windows" ]; then
            \cp ../conf/win/* ${hardware_path}/
        fi

        # packages
        cd ./packages
        if [ "x${os}" = x"windows" ]; then
            zip -rq ${hardware_dir_name}.zip ${hardware_dir_name}
        else
            tar -zcvf ${hardware_dir_name}.tar.gz ${hardware_dir_name}
        fi
        cd ..
        rm -rf ${hardware_path}
    done
done

\cp ./packages/* ./

if [ -d "./packages" ];then
    \rm -rf ./packages
fi

