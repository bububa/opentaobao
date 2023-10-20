package icburfq

import (
	"sync"
)

// RfqBuyRequestSearchDetailRemoteDto 结构体
type RfqBuyRequestSearchDetailRemoteDto struct {
	// 附件列表
	Attachments []Attachedfiles `json:"attachments,omitempty" xml:"attachments>attachedfiles,omitempty"`
	// RFQ详情
	RfqDetailDto *BuyRequestSearchDetailRemoteDto `json:"rfq_detail_dto,omitempty" xml:"rfq_detail_dto,omitempty"`
}

var poolRfqBuyRequestSearchDetailRemoteDto = sync.Pool{
	New: func() any {
		return new(RfqBuyRequestSearchDetailRemoteDto)
	},
}

// GetRfqBuyRequestSearchDetailRemoteDto() 从对象池中获取RfqBuyRequestSearchDetailRemoteDto
func GetRfqBuyRequestSearchDetailRemoteDto() *RfqBuyRequestSearchDetailRemoteDto {
	return poolRfqBuyRequestSearchDetailRemoteDto.Get().(*RfqBuyRequestSearchDetailRemoteDto)
}

// ReleaseRfqBuyRequestSearchDetailRemoteDto 释放RfqBuyRequestSearchDetailRemoteDto
func ReleaseRfqBuyRequestSearchDetailRemoteDto(v *RfqBuyRequestSearchDetailRemoteDto) {
	v.Attachments = v.Attachments[:0]
	v.RfqDetailDto = nil
	poolRfqBuyRequestSearchDetailRemoteDto.Put(v)
}
