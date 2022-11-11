## Beego Base Api Project

### 下载
```
git clone https://github.com/crayonxiaoxin/Beego_Base_Api.git
```

### 准备
安装 bee 工具，用于生成路由和文档 ([参考资料](http://beego.gocn.vip/beego/zh/developing/))
```
go get -u github.com/beego/beego/v2@latest
go get -u github.com/beego/bee/v2
``` 

### 运行
```
go mod tidy

bee run
```

### 生成路由
```
bee generate routers
```

### 生成文档
```
bee generate docs
```
文档地址：/swagger


### 目录结构
```
.
├── README.md
├── conf
│   └── app.conf                // 配置文件
├── controllers                 // 控制器
│   ├── login.go
│   ├── register.go
│   └── user.go
├── go.mod
├── go.sum
├── main.go
├── models                      // 模型
│   ├── login.go
│   └── user.go
├── routers                     // 路由
│   ├── commentsRouter.go       // bee generate routers 生成
│   └── router.go               // 路由配置
├── swagger                     // api 文档
│   ├── swagger.json            // bee generate docs 生成
│   └── swagger.yml
├── tests
└── utils
    ├── db.go                   // 连接数据库
    ├── gorm.go                 // 自定义 GORM Model
    ├── password.go             // 密码加密解密
    ├── rc.go                   // 统一返回
    └── token.go                // jwt token
```