package tmallservice

import (
	"sync"
)

// WorkcardBillInfoDto 结构体
type WorkcardBillInfoDto struct {
	// 费用明细
	Details []Double `json:"details,omitempty" xml:"details>double,omitempty"`
	// 工单ID
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
	// 分成金额（分）
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 平台抽佣比例 0.0 ~ 1.0
	PlatformCommissionRate int64 `json:"platform_commission_rate,omitempty" xml:"platform_commission_rate,omitempty"`
}

var poolWorkcardBillInfoDto = sync.Pool{
	New: func() any {
		return new(WorkcardBillInfoDto)
	},
}

// GetWorkcardBillInfoDto() 从对象池中获取WorkcardBillInfoDto
func GetWorkcardBillInfoDto() *WorkcardBillInfoDto {
	return poolWorkcardBillInfoDto.Get().(*WorkcardBillInfoDto)
}

// ReleaseWorkcardBillInfoDto 释放WorkcardBillInfoDto
func ReleaseWorkcardBillInfoDto(v *WorkcardBillInfoDto) {
	v.Details = v.Details[:0]
	v.WorkcardId = 0
	v.Amount = 0
	v.PlatformCommissionRate = 0
	poolWorkcardBillInfoDto.Put(v)
}
