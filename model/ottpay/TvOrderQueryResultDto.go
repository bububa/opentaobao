package ottpay

import (
	"sync"
)

// TvOrderQueryResultDto 结构体
type TvOrderQueryResultDto struct {
	// cpOrderNo
	CpOrderNo string `json:"cp_order_no,omitempty" xml:"cp_order_no,omitempty"`
	// orderNo
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// price
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// status
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// statusDesc
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
}

var poolTvOrderQueryResultDto = sync.Pool{
	New: func() any {
		return new(TvOrderQueryResultDto)
	},
}

// GetTvOrderQueryResultDto() 从对象池中获取TvOrderQueryResultDto
func GetTvOrderQueryResultDto() *TvOrderQueryResultDto {
	return poolTvOrderQueryResultDto.Get().(*TvOrderQueryResultDto)
}

// ReleaseTvOrderQueryResultDto 释放TvOrderQueryResultDto
func ReleaseTvOrderQueryResultDto(v *TvOrderQueryResultDto) {
	v.CpOrderNo = ""
	v.OrderNo = ""
	v.Price = ""
	v.Status = ""
	v.StatusDesc = ""
	poolTvOrderQueryResultDto.Put(v)
}
