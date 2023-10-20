package servicecenter

import (
	"sync"
)

// OfnPreRedPacketFundRecordDto 结构体
type OfnPreRedPacketFundRecordDto struct {
	// 资产编号
	FundId int64 `json:"fund_id,omitempty" xml:"fund_id,omitempty"`
	// 变化金额
	ChangeAmount int64 `json:"change_amount,omitempty" xml:"change_amount,omitempty"`
}

var poolOfnPreRedPacketFundRecordDto = sync.Pool{
	New: func() any {
		return new(OfnPreRedPacketFundRecordDto)
	},
}

// GetOfnPreRedPacketFundRecordDto() 从对象池中获取OfnPreRedPacketFundRecordDto
func GetOfnPreRedPacketFundRecordDto() *OfnPreRedPacketFundRecordDto {
	return poolOfnPreRedPacketFundRecordDto.Get().(*OfnPreRedPacketFundRecordDto)
}

// ReleaseOfnPreRedPacketFundRecordDto 释放OfnPreRedPacketFundRecordDto
func ReleaseOfnPreRedPacketFundRecordDto(v *OfnPreRedPacketFundRecordDto) {
	v.FundId = 0
	v.ChangeAmount = 0
	poolOfnPreRedPacketFundRecordDto.Put(v)
}
