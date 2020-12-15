package controllers

import (
	//"encoding/json"
	"fmt"
	"github.com/bai615/Go-000/Week04/goblog/app/models/user"
	"github.com/bai615/Go-000/Week04/goblog/app/requests"
	"github.com/bai615/Go-000/Week04/goblog/pkg/auth"
	"github.com/bai615/Go-000/Week04/goblog/pkg/flash"
	"github.com/bai615/Go-000/Week04/goblog/pkg/view"
	//"github.com/thedevsaddam/govalidator"
	"net/http"
)

// AuthController 处理静态页面
type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

/*type userForm struct {
	Name            string `valid:"name"`
	Email           string `valid:"email"`
	Password        string `valid:"password"`
	PasswordComfirm string `valid:"password_comfirm"`
}*/

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {

	// 1. 初始化数据
	/*_user := userForm{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordComfirm: r.PostFormValue("password_comfirm"),
	}*/
	// 1. 初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordComfirm: r.PostFormValue("password_comfirm"),
	}
/*
	// 2. 表单规则
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20"},
		"email":            []string{"required", "min:4", "max:30", "email"},
		"password":         []string{"required", "min:6"},
		"password_comfirm": []string{"required"},
	}

	// 3. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
		},
		"password_comfirm": []string{
			"required:确认密码框为必填项",
		},
	}

	// 3. 配置选项
	opts := govalidator.Options{
		Data:          &_user,
		Rules:         rules,
		TagIdentifier: "valid", // Struct 标签标识符
		Messages:      messages,
	}

	// 4. 开始认证
	errs := govalidator.New(opts).ValidateStruct()
*/

    // 2. 表单规则
    errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		// 3. 表单不通过 —— 重新显示表单
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
		/*
			// 4.1 有错误发生，打印数据
			data, _ := json.MarshalIndent(errs, "", "  ")
			fmt.Fprint(w, string(data))
		*/
	} else {
		// 4. 验证成功，创建数据

		_user.Create()

		if _user.ID > 0 {
			// 登录用户并跳转到首页
			flash.Success("恭喜您注册成功！")
			auth.Login(_user)
		    //fmt.Fprint(w, "插入成功，ID 为"+_user.GetStringID())
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
		    w.WriteHeader(http.StatusInternalServerError)
		    fmt.Fprint(w, "注册失败，请联系管理员")
		}
	}

/*
	// 0. 初始化变量
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// 1. 表单验证
	// 2. 验证通过 —— 入库，并跳转到首页
	_user := user.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	_user.Create()

	if _user.ID > 0 {
		fmt.Fprint(w, "插入成功，ID 为"+_user.GetStringID())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "创建用户失败，请联系管理员")
	}

	// 3. 表单不通过 —— 重新显示表单
 */
}

// Login 显示登录表单
func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {

	// 新增会话数据
	//session.Put("uid", "1")
	// 读取会话数据#
	//fmt.Fprint(w, session.Get("uid"))
	// 删除会话数据#
	//session.Forget("uid")
	// 销毁整个会话
	//session.Flush()

	view.RenderSimple(w, view.D{}, "auth.login")
}

// DoLogin 处理登录表单提交
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化表单数据
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// 2. 尝试登录
	if err := auth.Attempt(email, password); err == nil {
		// 登录成功
		flash.Success("欢迎回来！")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// 3. 失败，显示错误提示
		view.RenderSimple(w, view.D{
			"Error":    err.Error(),
			"Email":    email,
			"Password": password,
		}, "auth.login")
	}
}

// Logout 退出登录
func (*AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout()
	flash.Success("您已退出登录")
	http.Redirect(w, r, "/", http.StatusFound)
}