package wdk

import (
	"sync"
)

// ChannelRefundDto 结构体
type ChannelRefundDto struct {
	// 退款渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 渠道对应的退款金额(单位分)
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}

var poolChannelRefundDto = sync.Pool{
	New: func() any {
		return new(ChannelRefundDto)
	},
}

// GetChannelRefundDto() 从对象池中获取ChannelRefundDto
func GetChannelRefundDto() *ChannelRefundDto {
	return poolChannelRefundDto.Get().(*ChannelRefundDto)
}

// ReleaseChannelRefundDto 释放ChannelRefundDto
func ReleaseChannelRefundDto(v *ChannelRefundDto) {
	v.ChannelCode = ""
	v.RefundAmount = 0
	poolChannelRefundDto.Put(v)
}
