package servicecenter

import (
	"sync"
)

// ArticleViewResult 结构体
type ArticleViewResult struct {
	// sku详情列表
	ArticleItemViewUnits []ArticleItemViewUnit `json:"article_item_view_units,omitempty" xml:"article_item_view_units>article_item_view_unit,omitempty"`
	// 服务code
	ArticleCode string `json:"article_code,omitempty" xml:"article_code,omitempty"`
	// 服务简介
	ArticleCommment string `json:"article_commment,omitempty" xml:"article_commment,omitempty"`
	// 服务名称
	ArticleName string `json:"article_name,omitempty" xml:"article_name,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 用户淘宝nick
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 服务图片地址
	PictUrl string `json:"pict_url,omitempty" xml:"pict_url,omitempty"`
}

var poolArticleViewResult = sync.Pool{
	New: func() any {
		return new(ArticleViewResult)
	},
}

// GetArticleViewResult() 从对象池中获取ArticleViewResult
func GetArticleViewResult() *ArticleViewResult {
	return poolArticleViewResult.Get().(*ArticleViewResult)
}

// ReleaseArticleViewResult 释放ArticleViewResult
func ReleaseArticleViewResult(v *ArticleViewResult) {
	v.ArticleItemViewUnits = v.ArticleItemViewUnits[:0]
	v.ArticleCode = ""
	v.ArticleCommment = ""
	v.ArticleName = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Nick = ""
	v.PictUrl = ""
	poolArticleViewResult.Put(v)
}
