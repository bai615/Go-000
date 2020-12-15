package middlewares

import (
	"github.com/bai615/Go-000/Week04/goblog/pkg/auth"
	"github.com/bai615/Go-000/Week04/goblog/pkg/flash"
	"net/http"
)

// Auth 登录用户才可访问
func Auth(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !auth.Check() {
			flash.Warning("登录用户才能访问此页面")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next(w, r)
	}
}
