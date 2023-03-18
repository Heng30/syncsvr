#### syncsvr
一个基于`Go gin`的http配置备份服务器。通过http `Post` 请求将数据存储到sqlite数据库中。通过http `Get` 请求获取数据。具体使用方法参考下面说明。

### 运行
- `make gen_denpendence`
- `make download-dependence`
- `make run`

### 使用
- `syncsvr -query=true` 获取`访问token`
- `syncsvr -add=<token>` 添加`访问token`
- `syncsvr -del=<token>` 删除`访问token`
- `syncsvr -upd=<old_token,new_token>` 更新`访问token`

在程序启动后会为每个`访问token`生成一张表。
- 通过`curl -X POST -d 'I am value' "http://localhost:8000/testToken/key"`。将`I am value` 存储到`testToken`表的`key`字段中。
- 通过`curl -v "http://localhost:8000/testToken/key"`。将`testToken`表中的`key`字段关联的值获取出来。

### 启用https
- 修改配置文件`EnableTLS: true`
- 程序启动会生成默认的证书和私钥文件，放置在`~/.config/syncsvr/cert.pem` 和 `~/.config/syncsvr/key.pem`。 默认证书仅支持本地回环地址访问。需要支持外部网络访问需要替换为自己的证书和私钥。

### 参考
- [go安装依赖包（go get, go module）](https://blog.csdn.net/weixin_41519463/article/details/103501485)
- [Golang设置代理](https://developer.aliyun.com/article/879662)
- [Gin 解决跨域问题跨域配置](https://juejin.cn/post/6871583587062415367)
- [Go 入门指南](https://learnku.com/docs/the-way-to-go)
