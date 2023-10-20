package tmallnr

import (
	"sync"
)

// PageData 结构体
type PageData struct {
	// 电子凭证实例DTO
	DataList []NrtCertificateInstanceResponseDto `json:"data_list,omitempty" xml:"data_list>nrt_certificate_instance_response_dto,omitempty"`
	// 显示数量
	PageSize string `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	CurrentPage string `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总数
	TotalItem string `json:"total_item,omitempty" xml:"total_item,omitempty"`
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
	v.PageSize = ""
	v.CurrentPage = ""
	v.TotalItem = ""
	poolPageData.Put(v)
}
