package tmallgenie

import (
	"sync"
)

// PurchaseCircleInfoForOuterDto 结构体
type PurchaseCircleInfoForOuterDto struct {
	// 圈选内容
	CircleContent string `json:"circle_content,omitempty" xml:"circle_content,omitempty"`
	// 圈选标识id,由采销系统分配
	CircleId int64 `json:"circle_id,omitempty" xml:"circle_id,omitempty"`
}

var poolPurchaseCircleInfoForOuterDto = sync.Pool{
	New: func() any {
		return new(PurchaseCircleInfoForOuterDto)
	},
}

// GetPurchaseCircleInfoForOuterDto() 从对象池中获取PurchaseCircleInfoForOuterDto
func GetPurchaseCircleInfoForOuterDto() *PurchaseCircleInfoForOuterDto {
	return poolPurchaseCircleInfoForOuterDto.Get().(*PurchaseCircleInfoForOuterDto)
}

// ReleasePurchaseCircleInfoForOuterDto 释放PurchaseCircleInfoForOuterDto
func ReleasePurchaseCircleInfoForOuterDto(v *PurchaseCircleInfoForOuterDto) {
	v.CircleContent = ""
	v.CircleId = 0
	poolPurchaseCircleInfoForOuterDto.Put(v)
}
