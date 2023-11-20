[中文文档](./README.zh-CN.md)

#### syncsvr
An http configuration backup server based on `Go gin`. Data is stored in an SQLite database through an http `Post` request. Data can be retrieved through an http `Get` request. Please refer to the following instructions for specific usage.

#### Running
- `make gen_denpendence`
- `make download-dependence`
- `make run`

#### Usage
- `syncsvr -query=true` to get the access token
- `syncsvr -add=<token>` to add an access token
- `syncsvr -del=<token>` to delete an access token
- `syncsvr -upd=<old_token,new_token>` to update an access token

After the program starts, a table will be generated for each access token.
- Use `curl -X POST -d 'I am value' "http://localhost:8000/testToken/key"` to store `'I am value'` in the `'key'` field of the `'testToken'` table.
- Use `curl -v "http://localhost:8000/testToken/key"` to retrieve the value associated with the `'key'` field in the `'testToken'` table.

#### Enable HTTPS
- Modify the configuration file: set `EnableTLS: true`
- The program will generate default certificate and private key files upon startup, placed in `~/.config/syncsvr/cert.pem` and `~/.config/syncsvr/key.pem`. The default certificate only supports accessing localhost. To support external network access, replace them with your own certificate and private key.

#### References
- [Installing dependency packages in Go (go get, go module)](https://blog.csdn.net/weixin_41519463/article/details/103501485)
- [Setting up a proxy in Golang](https://developer.aliyun.com/article/879662)
- [Gin: Solving Cross-Origin Resource Sharing (CORS) Issues](https://juejin.cn/post/6871583587062415367)
- [The Way to Go: A Thorough Introduction to the Go Programming Language](https://learnku.com/docs/the-way-to-go)
