# compile for version
make
if [ $? -ne 0 ]; then
    echo "make error"
    exit 1
fi

hardware_version=`./bin/hardware -v`
echo "build version: $cyssh_version"

# cross_compiles
make -f ./Makefile.cross-compiles

rm -rf ./dist/packages
mkdir -p ./dist/packages

os_all='linux darwin freebsd'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle riscv64'

cd ./dist

for os in $os_all; do
    for arch in $arch_all; do
        hardware_dir_name="hardware_${hardware_version}_${os}_${arch}"
        hardware_path="./packages/hardware_${hardware_version}_${os}_${arch}"
        
        if [ ! -f "./hardware_${os}_${arch}" ]; then
            continue
        fi
        mkdir ${hardware_path}
        mv ./cyssh_${os}_${arch} ${hardware_path}/cyssh
        mv ./cyscp_${os}_${arch} ${hardware_path}/cyscp
        mv ./cyssh-server_${os}_${arch} ${hardware_path}/cysh-server
        cp ../LICENSE ${hardware_path}
        #cp -rf ../conf/* ${hardware_path}

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

cd -
