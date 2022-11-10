## Beego Base Api Project

#### 运行

`go mod tidy`

`bee run`


#### 目录结构

```
.
├── README.md
├── conf
│   └── app.conf
├── controllers
│   ├── login.go
│   ├── register.go
│   └── user.go
├── go.mod
├── go.sum
├── main.go
├── models
│   ├── login.go
│   └── user.go
├── routers
│   ├── commentsRouter.go
│   └── router.go
├── swagger
│   ├── swagger.json
│   └── swagger.yml
├── tests
└── utils
    ├── db.go
    ├── gorm.go
    ├── password.go
    ├── rc.go
    └── token.go
```