#!/bin/bash
# 获取sh文件的绝对文件
CURDIR=$(cd $(dirname ${BASH_SOURCE[0]}); pwd )
# 尝试创建build文件夹
mkdir ./build/;
# 尝试清空build文件夹
rm -rf ./build/*;
# 尝试创建二进制文件夹
mkdir ./build/bin;
echo "创建bin文件夹完成"
# 复制pandoc到打包目录
cp ./bin/pandoc ./build/bin/
echo "复制pandoc完成"
# 复制disable_float到打包目录
cp ./bin/disable_float.tex ./build/bin/
echo "复制disable_float完成"
# 复制配置文件
cp ./config.json ./build
echo "复制配置文件完成"
# 复制静态文件夹
cp -r ./assets/ build/
echo "制静态文件夹完成"
#使用go-assets-builder编译二进制文件
./bin/go-assets-builder ./html/ -o ./bindata/asset.go -p bindata
echo "go-assets-builder编译完成"
# 打包
go build -o ./build/main main.go
echo "打包完成"
# 如果执行的是build则不执行程序
if [ "$1" != "build" ]
then
  # shellcheck disable=SC2164
  cd ./build/
  echo "进入build"
  ./main
else
  echo "打包成功，打包地址：$CURDIR/build"
fi