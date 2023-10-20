package tanx

import (
	"sync"
)

// BiddingRefuseDto 结构体
type BiddingRefuseDto struct {
	// 创意级别对应的错误码
	FilterId string `json:"filter_id,omitempty" xml:"filter_id,omitempty"`
	// Pv粒度错误码对应描述二级原因
	FilterIdDesc string `json:"filter_id_desc,omitempty" xml:"filter_id_desc,omitempty"`
	// Pv粒度错误码对应的一级原因
	FilterClassDesc string `json:"filter_class_desc,omitempty" xml:"filter_class_desc,omitempty"`
	// O123545
	CreativeId string `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 创意过滤次数
	AdfilPv int64 `json:"adfil_pv,omitempty" xml:"adfil_pv,omitempty"`
	// dsp_id
	DspId int64 `json:"dsp_id,omitempty" xml:"dsp_id,omitempty"`
}

var poolBiddingRefuseDto = sync.Pool{
	New: func() any {
		return new(BiddingRefuseDto)
	},
}

// GetBiddingRefuseDto() 从对象池中获取BiddingRefuseDto
func GetBiddingRefuseDto() *BiddingRefuseDto {
	return poolBiddingRefuseDto.Get().(*BiddingRefuseDto)
}

// ReleaseBiddingRefuseDto 释放BiddingRefuseDto
func ReleaseBiddingRefuseDto(v *BiddingRefuseDto) {
	v.FilterId = ""
	v.FilterIdDesc = ""
	v.FilterClassDesc = ""
	v.CreativeId = ""
	v.AdfilPv = 0
	v.DspId = 0
	poolBiddingRefuseDto.Put(v)
}
