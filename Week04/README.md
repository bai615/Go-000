Week04 作业题目：

1. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

以上作业，要求提交到 GitHub 上面，Week04 作业提交地址：
https://github.com/Go-000/Go-000/issues/76



## 目录结构

```
├── api
│   └── article
│       ├── article.pb.go
│       └── article.proto
├── app
│   ├── http
│   │   ├── controllers
│   │   │   ├── articles_controller.go
│   │   │   ├── auth_controller.go
│   │   │   ├── base_controller.go
│   │   │   ├── pages_controller.go
│   │   │   └── user_controller.go
│   │   └── middlewares
│   │       ├── auth.go
│   │       ├── force_html.go
│   │       ├── guest.go
│   │       ├── middleware.go
│   │       ├── remove_trailing_slash.go
│   │       └── start_session.go
│   ├── models
│   │   ├── article
│   │   │   ├── article.go
│   │   │   └── crud.go
│   │   ├── model.go
│   │   └── user
│   │       ├── crud.go
│   │       ├── hooks.go
│   │       └── user.go
│   ├── policies
│   │   └── topic_policy.go
│   └── requests
│       ├── article_form.go
│       ├── request.go
│       └── user_registration.go
├── bootstrap
│   ├── db.go
│   └── route.go
├── config
│   ├── app.go
│   ├── config.go
│   ├── database.go
│   └── session.go
├── go.mod
├── go.sum
├── main.go
├── pkg
│   ├── auth
│   │   └── auth.go
│   ├── config
│   │   └── config.go
│   ├── database
│   │   └── database.go
│   ├── flash
│   │   └── flash.go
│   ├── logger
│   │   └── logger.go
│   ├── model
│   │   └── model.go
│   ├── pagination
│   │   └── pagination.go
│   ├── password
│   │   └── password.go
│   ├── route
│   │   └── router.go
│   ├── session
│   │   └── session.go
│   ├── types
│   │   └── converter.go
│   └── view
│       └── view.go
├── public
│   ├── css
│   │   ├── app.css
│   │   └── bootstrap.min.css
│   └── js
│       └── bootstrap.min.js
├── resources
│   └── views
│       ├── articles
│       │   ├── _article_meta.gohtml
│       │   ├── _form_field.gohtml
│       │   ├── create.gohtml
│       │   ├── edit.gohtml
│       │   ├── index.gohtml
│       │   └── show.gohtml
│       ├── auth
│       │   ├── login.gohtml
│       │   └── register.gohtml
│       └── layouts
│           ├── _form_error_feedback.gohtml
│           ├── _messages.gohtml
│           ├── _pagination.gohtml
│           ├── app.gohtml
│           ├── sidebar.gohtml
│           └── simple.gohtml
├── routes
│   └── web.go
├── tests
│   └── pages_test.go
└── tmp

```
