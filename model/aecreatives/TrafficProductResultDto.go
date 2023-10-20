package aecreatives

import (
	"sync"
)

// TrafficProductResultDto 结构体
type TrafficProductResultDto struct {
	// 返回商品列表
	Products []Product `json:"products,omitempty" xml:"products>product,omitempty"`
	// 当前返回页数
	CurrentPageNo int64 `json:"current_page_no,omitempty" xml:"current_page_no,omitempty"`
	// 当前返回数量
	CurrentRecordCount int64 `json:"current_record_count,omitempty" xml:"current_record_count,omitempty"`
	// 总计页数
	TotalPageNo int64 `json:"total_page_no,omitempty" xml:"total_page_no,omitempty"`
	// 总计数量
	TotalRecordCount int64 `json:"total_record_count,omitempty" xml:"total_record_count,omitempty"`
	// 数据是否拉取完成
	IsFinished bool `json:"is_finished,omitempty" xml:"is_finished,omitempty"`
}

var poolTrafficProductResultDto = sync.Pool{
	New: func() any {
		return new(TrafficProductResultDto)
	},
}

// GetTrafficProductResultDto() 从对象池中获取TrafficProductResultDto
func GetTrafficProductResultDto() *TrafficProductResultDto {
	return poolTrafficProductResultDto.Get().(*TrafficProductResultDto)
}

// ReleaseTrafficProductResultDto 释放TrafficProductResultDto
func ReleaseTrafficProductResultDto(v *TrafficProductResultDto) {
	v.Products = v.Products[:0]
	v.CurrentPageNo = 0
	v.CurrentRecordCount = 0
	v.TotalPageNo = 0
	v.TotalRecordCount = 0
	v.IsFinished = false
	poolTrafficProductResultDto.Put(v)
}
