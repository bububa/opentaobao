package drugtrace

import (
	"sync"
)

// Page 结构体
type Page struct {
	// 返回列表
	ResultList []PEntParDto `json:"result_list,omitempty" xml:"result_list>p_ent_par_dto,omitempty"`
	// 总数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 当前页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPage = sync.Pool{
	New: func() any {
		return new(Page)
	},
}

// GetPage() 从对象池中获取Page
func GetPage() *Page {
	return poolPage.Get().(*Page)
}

// ReleasePage 释放Page
func ReleasePage(v *Page) {
	v.ResultList = v.ResultList[:0]
	v.TotalNum = 0
	v.Page = 0
	v.PageSize = 0
	poolPage.Put(v)
}
