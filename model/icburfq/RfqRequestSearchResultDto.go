package icburfq

import (
	"sync"
)

// RfqRequestSearchResultDto 结构体
type RfqRequestSearchResultDto struct {
	// RFQ列表
	RequestList []Requestlist `json:"request_list,omitempty" xml:"request_list>requestlist,omitempty"`
	// 类目列表
	CategoryList []Categorylist `json:"category_list,omitempty" xml:"category_list>categorylist,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolRfqRequestSearchResultDto = sync.Pool{
	New: func() any {
		return new(RfqRequestSearchResultDto)
	},
}

// GetRfqRequestSearchResultDto() 从对象池中获取RfqRequestSearchResultDto
func GetRfqRequestSearchResultDto() *RfqRequestSearchResultDto {
	return poolRfqRequestSearchResultDto.Get().(*RfqRequestSearchResultDto)
}

// ReleaseRfqRequestSearchResultDto 释放RfqRequestSearchResultDto
func ReleaseRfqRequestSearchResultDto(v *RfqRequestSearchResultDto) {
	v.RequestList = v.RequestList[:0]
	v.CategoryList = v.CategoryList[:0]
	v.Total = 0
	poolRfqRequestSearchResultDto.Put(v)
}
