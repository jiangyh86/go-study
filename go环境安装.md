# Go环境安装

## 一、安装下载

[环境安装下载地址](https://golang.google.cn/dl/)

![image-20221103123003857](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221103123003857.png)

[1.19.2oss下载地址](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/go1.19.2.windows-amd64.msi)

![image-20221103123202550](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221103123202550.png)

直接下一步就可以

## 二、版本查看

安装完成后，会自动添加到环境变量path

输入 go version 可以看到版本

![image-20221103123317434](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221103123317434.png)

## 三、设置GOROOT和GOPATH

1. 创建GoWorks文件夹以及子目录

![image-20221103123617302](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221103123617302.png)

![image-20221103123637698](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221103123637698.png)

2. 设置环境变量

![image-20221103123711744](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221103123711744.png)

## 四、查看环境是否配置完成

1. cmd 输入 go env 查看当前go的安装环境信息

![image-20221103124019131](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221103124019131.png) 

与设置的一样即可，如果不一样树C盘的默认目录

==在用户变量下默认创建的GOROOT和GOPATH改成环境变量的值即可==

![image-20221103124124967](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221103124124967.png)