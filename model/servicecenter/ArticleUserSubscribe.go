package servicecenter

import (
	"sync"
)

// ArticleUserSubscribe 结构体
type ArticleUserSubscribe struct {
	// 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 订购关系到期时间
	Deadline string `json:"deadline,omitempty" xml:"deadline,omitempty"`
}

var poolArticleUserSubscribe = sync.Pool{
	New: func() any {
		return new(ArticleUserSubscribe)
	},
}

// GetArticleUserSubscribe() 从对象池中获取ArticleUserSubscribe
func GetArticleUserSubscribe() *ArticleUserSubscribe {
	return poolArticleUserSubscribe.Get().(*ArticleUserSubscribe)
}

// ReleaseArticleUserSubscribe 释放ArticleUserSubscribe
func ReleaseArticleUserSubscribe(v *ArticleUserSubscribe) {
	v.ItemCode = ""
	v.Deadline = ""
	poolArticleUserSubscribe.Put(v)
}
