package servicecenter

import (
	"sync"
)

// ExternalEndRedPacketDto 结构体
type ExternalEndRedPacketDto struct {
	// 红包支付状态
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 红包状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 发放总金额，单位：分
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 已打款金额，单位：分
	DisburseFee int64 `json:"disburse_fee,omitempty" xml:"disburse_fee,omitempty"`
	// 退回金额，单位：分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 已冻结金额，单位：分
	FreezeFee int64 `json:"freeze_fee,omitempty" xml:"freeze_fee,omitempty"`
	// 可用金额，单位：分
	AvailableFee int64 `json:"available_fee,omitempty" xml:"available_fee,omitempty"`
}

var poolExternalEndRedPacketDto = sync.Pool{
	New: func() any {
		return new(ExternalEndRedPacketDto)
	},
}

// GetExternalEndRedPacketDto() 从对象池中获取ExternalEndRedPacketDto
func GetExternalEndRedPacketDto() *ExternalEndRedPacketDto {
	return poolExternalEndRedPacketDto.Get().(*ExternalEndRedPacketDto)
}

// ReleaseExternalEndRedPacketDto 释放ExternalEndRedPacketDto
func ReleaseExternalEndRedPacketDto(v *ExternalEndRedPacketDto) {
	v.PayStatus = ""
	v.Status = 0
	v.TotalFee = 0
	v.DisburseFee = 0
	v.RefundFee = 0
	v.FreezeFee = 0
	v.AvailableFee = 0
	poolExternalEndRedPacketDto.Put(v)
}
