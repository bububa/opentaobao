package alihealth2

import (
	"sync"
)

// PageInfoDto 结构体
type PageInfoDto struct {
	// result
	Results []AlibabaAlihealthTracecodesellerBillResultSearchResult `json:"results,omitempty" xml:"results>alibaba_alihealth_tracecodeseller_bill_result_search_result,omitempty"`
	// 总数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// page
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// pageSize
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPageInfoDto = sync.Pool{
	New: func() any {
		return new(PageInfoDto)
	},
}

// GetPageInfoDto() 从对象池中获取PageInfoDto
func GetPageInfoDto() *PageInfoDto {
	return poolPageInfoDto.Get().(*PageInfoDto)
}

// ReleasePageInfoDto 释放PageInfoDto
func ReleasePageInfoDto(v *PageInfoDto) {
	v.Results = v.Results[:0]
	v.TotalNum = 0
	v.Page = 0
	v.PageSize = 0
	poolPageInfoDto.Put(v)
}
