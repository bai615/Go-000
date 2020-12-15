package policies

import (
	"github.com/bai615/Go-000/Week04/goblog/app/models/article"
	"github.com/bai615/Go-000/Week04/goblog/pkg/auth"
)

// CanModifyArticle 是否允许修改话题
func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}