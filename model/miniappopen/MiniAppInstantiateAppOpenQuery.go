package miniappopen

import (
	"sync"
)

// MiniAppInstantiateAppOpenQuery 结构体
type MiniAppInstantiateAppOpenQuery struct {
	// 模版id
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

var poolMiniAppInstantiateAppOpenQuery = sync.Pool{
	New: func() any {
		return new(MiniAppInstantiateAppOpenQuery)
	},
}

// GetMiniAppInstantiateAppOpenQuery() 从对象池中获取MiniAppInstantiateAppOpenQuery
func GetMiniAppInstantiateAppOpenQuery() *MiniAppInstantiateAppOpenQuery {
	return poolMiniAppInstantiateAppOpenQuery.Get().(*MiniAppInstantiateAppOpenQuery)
}

// ReleaseMiniAppInstantiateAppOpenQuery 释放MiniAppInstantiateAppOpenQuery
func ReleaseMiniAppInstantiateAppOpenQuery(v *MiniAppInstantiateAppOpenQuery) {
	v.TemplateId = ""
	v.PageSize = 0
	v.PageNum = 0
	poolMiniAppInstantiateAppOpenQuery.Put(v)
}
