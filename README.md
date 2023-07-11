# webdav
a golang webdav server!

```
[root@ webdav]# go build
[root@ webdav]# ./webdav -i /fdjbjdbvd/ -d "app,/xieyuhua;dav1,/www/wwwroot/"

[root@ webdav]# ./webdav -h
Usage of ./webdav:
  -d string
    	目录 'app,/app;abc,/abc' (default "app,./")
  -i string
    	定义随机入口,保护入口安全, 例如：/abc/ (default "/")
  -p string
    	监听服务和端口 (default "7777")
[root@ webdav]# 
```