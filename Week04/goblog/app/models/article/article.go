package article

import (
	"github.com/bai615/Go-000/Week04/goblog/app/models"
	"github.com/bai615/Go-000/Week04/goblog/app/models/user"
	"github.com/bai615/Go-000/Week04/goblog/pkg/route"
	"github.com/bai615/Go-000/Week04/goblog/pkg/types"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`

	UserID uint64 `gorm:"not null;index"`
	User   user.User
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

// GetStringID 获取 ID 的字符串格式
func (a Article) GetStringID() string {
	return types.Uint64ToString(a.ID)
}

// CreatedAtDate 创建日期
func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}