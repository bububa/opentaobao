package tmallgenie

import (
	"sync"
)

// PurchaseSendPlanDto 结构体
type PurchaseSendPlanDto struct {
	// 发放活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 领取规则
	ReceiveRule string `json:"receive_rule,omitempty" xml:"receive_rule,omitempty"`
	// 商品信息
	PurchaseItemInfo string `json:"purchase_item_info,omitempty" xml:"purchase_item_info,omitempty"`
	// 活动id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolPurchaseSendPlanDto = sync.Pool{
	New: func() any {
		return new(PurchaseSendPlanDto)
	},
}

// GetPurchaseSendPlanDto() 从对象池中获取PurchaseSendPlanDto
func GetPurchaseSendPlanDto() *PurchaseSendPlanDto {
	return poolPurchaseSendPlanDto.Get().(*PurchaseSendPlanDto)
}

// ReleasePurchaseSendPlanDto 释放PurchaseSendPlanDto
func ReleasePurchaseSendPlanDto(v *PurchaseSendPlanDto) {
	v.Name = ""
	v.StartTime = ""
	v.EndTime = ""
	v.ReceiveRule = ""
	v.PurchaseItemInfo = ""
	v.Id = 0
	poolPurchaseSendPlanDto.Put(v)
}
