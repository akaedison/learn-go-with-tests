

git 源码安装

##### 1. 下载源码

````
wget https://github.com/git/git/archive/refs/tags/v2.31.0.tar.gz
````

##### 2. 解压

````
tar -zxvf v2.31.0.tar.gz
````

##### 3.编译安装

````
cd v2.31.0
.configure -prefix=/usr/lcoal/
````

报错: 提示没有 .configure

解决: 安装 autoconf 生成

````
yum install autoconf
autoconf
.configure -prefix=/usr/local/
````

报错:  no acceptable C compiler found in $PATH 没有 C 编译器

解决:  安装 C 编译器

````
yum install gcc
.configure -prefix=/usr/local/
make && make install
````

报错:  fatal error: zlib.h: No such file or directory 没有 zlib.h

解决: 安装zlib-devel

````
yum install zlib-devel
make && make install
````

出现

````
	rm -f "$execdir/$p" && \
	if test -z ""; \
	then \
		test -n "" && \
		ln -s "$destdir_from_execdir_SQ/bin/git" "$execdir/$p" || \
		{ test -z "" && \
		  ln "$execdir/git" "$execdir/$p" 2>/dev/null || \
		  ln -s "git" "$execdir/$p" 2>/dev/null || \
		  cp "$execdir/git" "$execdir/$p" || exit; }; \
	fi \
done && \
remote_curl_aliases="" && \
for p in $remote_curl_aliases; do \
	rm -f "$execdir/$p" && \
	test -n "" && \
	ln -s "git-remote-http" "$execdir/$p" || \
	{ test -z "" && \
	  ln "$execdir/git-remote-http" "$execdir/$p" 2>/dev/null || \
	  ln -s "git-remote-http" "$execdir/$p" 2>/dev/null || \
	  cp "$execdir/git-remote-http" "$execdir/$p" || exit; } \
done && \
./check_bindir "z$bindir" "z$execdir" "$bindir/git-add"

````

大概这么一串,说明成功,cd 到安装路径

````
cd bin 
./git
````

现在说明安装成功,但是还不能在别处运行,我们添加环境变量

##### 4.配置环境变量

永久配置环境变量有两种方法,一种是直接编辑 /etc/profile

````
vim /etc/profile
````

然后在最后追加

````
export PATH=$PATH:/usr/lcoal/soft/git/bin
````

直接编辑profile 文件是不好维护的,我采用了另外一种方法

````
cd /etc/profile.d
vim git.sh
````

我们在profile.d下新建一个脚本,profile配置文件中会读取这个文件夹下所有的sh,新建内容:

````
export GIT_HOME=/usr/local/soft/git
export PATH=$PATH:$GIT_HOME/bin
````

然后

````
source /etc/profile
````

生效,我们测试一次写入的环境变量

````
echo $GIT_HOME
/usr/local/soft/git
````

环境变量生效,现在测试git

````
git version
git version 2.31.0
````

git 安装成功,虽然yum install git 可以直接安装,但是版本太老,对centos 感兴趣的建议源码安装,从中遇到问题,解决问题,可以学到很多的.

