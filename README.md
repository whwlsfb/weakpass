## 使用方法
帮助信息：  
```
./pwd -h
```
```
Usage of ./pwd:
  -c string
    	请输入主要密码字段：公司名,多个字段，以空格隔开
  -d string
    	请输入时间：2019-01-01，多个字段，以空格隔开
  -e string
    	请输入email:xxxx@xx.com
  -n string
    	请输入主要密码字段：用户名,多个字段，以空格隔开
  -s string
    	请输入次要密码字段，自定义，多个字段，以空格隔开
```
## 生成linux及windows执行文件
```
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build pwd.go
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build pwd.go
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build pwd.go
$ CGO_ENABLED=0 GOOS=windows GOARCH=386 go build pwd.go
```

## 文件说明
```
pwd.go                    代码文件
result_weak_pass.txt      生成弱口令后的密码存放文件,会自动生成
weakpass.txt              收集默认弱口令
initpass.txt              默认密码最小字段
```
