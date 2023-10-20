package txcs

import (
	"sync"
)

// WebPageData 结构体
type WebPageData struct {
	// 结果数据
	List []WebPageDataList `json:"list,omitempty" xml:"list>web_page_data_list,omitempty"`
	// 页码信息
	Pagination *Pagination `json:"pagination,omitempty" xml:"pagination,omitempty"`
}

var poolWebPageData = sync.Pool{
	New: func() any {
		return new(WebPageData)
	},
}

// GetWebPageData() 从对象池中获取WebPageData
func GetWebPageData() *WebPageData {
	return poolWebPageData.Get().(*WebPageData)
}

// ReleaseWebPageData 释放WebPageData
func ReleaseWebPageData(v *WebPageData) {
	v.List = v.List[:0]
	v.Pagination = nil
	poolWebPageData.Put(v)
}
