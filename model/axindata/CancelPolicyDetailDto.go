package axindata

import (
	"sync"
)

// CancelPolicyDetailDto 结构体
type CancelPolicyDetailDto struct {
	// 扣除值
	DeductFeeRate int64 `json:"deduct_fee_rate,omitempty" xml:"deduct_fee_rate,omitempty"`
	// 提前小时数
	AheadCancelHours int64 `json:"ahead_cancel_hours,omitempty" xml:"ahead_cancel_hours,omitempty"`
}

var poolCancelPolicyDetailDto = sync.Pool{
	New: func() any {
		return new(CancelPolicyDetailDto)
	},
}

// GetCancelPolicyDetailDto() 从对象池中获取CancelPolicyDetailDto
func GetCancelPolicyDetailDto() *CancelPolicyDetailDto {
	return poolCancelPolicyDetailDto.Get().(*CancelPolicyDetailDto)
}

// ReleaseCancelPolicyDetailDto 释放CancelPolicyDetailDto
func ReleaseCancelPolicyDetailDto(v *CancelPolicyDetailDto) {
	v.DeductFeeRate = 0
	v.AheadCancelHours = 0
	poolCancelPolicyDetailDto.Put(v)
}
