package simba

import (
	"sync"
)

// KeywordPage 结构体
type KeywordPage struct {
	// 关键词列表
	KeywordList []Keyword `json:"keyword_list,omitempty" xml:"keyword_list>keyword,omitempty"`
	// 每页数据大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 返回第几页的数据从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 所查询的数据总数，只有当返回第一页数据时有值，当要求返回的数据非第一页时，此返回值无效
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
}

var poolKeywordPage = sync.Pool{
	New: func() any {
		return new(KeywordPage)
	},
}

// GetKeywordPage() 从对象池中获取KeywordPage
func GetKeywordPage() *KeywordPage {
	return poolKeywordPage.Get().(*KeywordPage)
}

// ReleaseKeywordPage 释放KeywordPage
func ReleaseKeywordPage(v *KeywordPage) {
	v.KeywordList = v.KeywordList[:0]
	v.PageSize = 0
	v.PageNo = 0
	v.TotalItem = 0
	poolKeywordPage.Put(v)
}
