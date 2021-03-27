#### docker 二进制安装

1.下载二进制压缩包

[稳定版](https://download.docker.com/linux/static/stable/)

````
wget https://download.docker.com/linux/static/stable/x86_64/docker-20.10.0.tgz
````

2.解压

````
tar zxvf docker-20.10.0.tgz
````

3.放进path

````
cp docker/* /usr/bin
````

4.开启 docker daemon

````
docker &
````

5.验证

````
docker run hello-world
````

6.加入service

二进制安装的docker 是不能直接用

````
systemctl start docker
````

等命令的,提示 docker.service找不到

那么直接加一个

https://github.com/moby/moby/blob/master/contrib/init/systemd/docker.service.rpm

新版的docker 只需要一个service 不需要socket

````
vim /etc/systemd/system/docker.service
````



````
[Unit]
Description=Docker Application Container Engine
Documentation=https://docs.docker.com
After=network-online.target firewalld.service
Wants=network-online.target

[Service]
Type=notify
# the default is not to use systemd for cgroups because the delegate issues still
# exists and systemd currently does not support the cgroup feature set required
# for containers run by docker
ExecStart=/usr/bin/dockerd
ExecReload=/bin/kill -s HUP $MAINPID
# Having non-zero Limit*s causes performance problems due to accounting overhead
# in the kernel. We recommend using cgroups to do container-local accounting.
LimitNOFILE=infinity
LimitNPROC=infinity
LimitCORE=infinity
# Uncomment TasksMax if your systemd version supports it.
# Only systemd 226 and above support this version.
#TasksMax=infinity
TimeoutStartSec=0
# set delegate yes so that systemd does not reset the cgroups of docker containers
Delegate=yes
# kill only the docker process, not all processes in the cgroup
KillMode=process
# restart the docker process if it exits prematurely
Restart=on-failure
StartLimitBurst=3
StartLimitInterval=60s

[Install]
WantedBy=multi-user.target
````



````
systemctl deamon-reload
````

即可生效





