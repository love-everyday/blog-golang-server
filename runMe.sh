 #!/bin/bash

if [ -z $1 ] ; then
 echo '请输入镜像版本'
 exit 1
fi
echo '镜像版本是： '$1
docker build -t blog:$1 . 
imageId=`docker images | grep blog | awk '{print $3}' | sed -n '1p'`
echo $imageId
docker tag $imageId registry.cn-hangzhou.aliyuncs.com/blog-yuan/blog:$1
docker push registry.cn-hangzhou.aliyuncs.com/blog-yuan/blog:$1
