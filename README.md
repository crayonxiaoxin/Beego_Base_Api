## Beego Base Api Project

### 准备
```
go get -u github.com/beego/beego/v2@latest
go get -u github.com/beego/bee/v2
``` 

### 运行
```
git clone https://github.com/crayonxiaoxin/Beego_Base_Api.git

cd Beego_Base_Api

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
文档路径：/swagger


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