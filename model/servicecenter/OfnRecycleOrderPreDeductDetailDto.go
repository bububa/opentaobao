package servicecenter

import (
	"sync"
)

// OfnRecycleOrderPreDeductDetailDto 结构体
type OfnRecycleOrderPreDeductDetailDto struct {
	// 数据更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 前置抵扣红包发放总金额，单位 分
	PreDeductTotalFee int64 `json:"pre_deduct_total_fee,omitempty" xml:"pre_deduct_total_fee,omitempty"`
	// 前置抵扣红包使用金额，单位 分
	PreDeductUsedFee int64 `json:"pre_deduct_used_fee,omitempty" xml:"pre_deduct_used_fee,omitempty"`
	// 前置抵扣红包退回金额，单位 分
	PreDeductRefundFee int64 `json:"pre_deduct_refund_fee,omitempty" xml:"pre_deduct_refund_fee,omitempty"`
	// 前置抵扣红包 ID
	PreDeductFundId int64 `json:"pre_deduct_fund_id,omitempty" xml:"pre_deduct_fund_id,omitempty"`
	// 线下已结算金额，单位 分【正数表示服务商已给用户的钱，为负表示用户已给服务商的钱】
	OfflineSettleFee int64 `json:"offline_settle_fee,omitempty" xml:"offline_settle_fee,omitempty"`
	// 质检价，单位 分
	QaAmount int64 `json:"qa_amount,omitempty" xml:"qa_amount,omitempty"`
}

var poolOfnRecycleOrderPreDeductDetailDto = sync.Pool{
	New: func() any {
		return new(OfnRecycleOrderPreDeductDetailDto)
	},
}

// GetOfnRecycleOrderPreDeductDetailDto() 从对象池中获取OfnRecycleOrderPreDeductDetailDto
func GetOfnRecycleOrderPreDeductDetailDto() *OfnRecycleOrderPreDeductDetailDto {
	return poolOfnRecycleOrderPreDeductDetailDto.Get().(*OfnRecycleOrderPreDeductDetailDto)
}

// ReleaseOfnRecycleOrderPreDeductDetailDto 释放OfnRecycleOrderPreDeductDetailDto
func ReleaseOfnRecycleOrderPreDeductDetailDto(v *OfnRecycleOrderPreDeductDetailDto) {
	v.UpdateTime = ""
	v.PreDeductTotalFee = 0
	v.PreDeductUsedFee = 0
	v.PreDeductRefundFee = 0
	v.PreDeductFundId = 0
	v.OfflineSettleFee = 0
	v.QaAmount = 0
	poolOfnRecycleOrderPreDeductDetailDto.Put(v)
}
