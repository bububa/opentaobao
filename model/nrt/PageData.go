package nrt

import (
	"sync"
)

// PageData 结构体
type PageData struct {
	// 数据列表
	DataList []CouponTemplateDto `json:"data_list,omitempty" xml:"data_list>coupon_template_dto,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 1
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPageData = sync.Pool{
	New: func() any {
		return new(PageData)
	},
}

// GetPageData() 从对象池中获取PageData
func GetPageData() *PageData {
	return poolPageData.Get().(*PageData)
}

// ReleasePageData 释放PageData
func ReleasePageData(v *PageData) {
	v.DataList = v.DataList[:0]
	v.CurrentPage = 0
	v.PageSize = 0
	poolPageData.Put(v)
}
