package alihouse

import (
	"sync"
)

// PreDepositGoldUnbindDto 结构体
type PreDepositGoldUnbindDto struct {
	// 外部活动id
	OuterSalesActivityId string `json:"outer_sales_activity_id,omitempty" xml:"outer_sales_activity_id,omitempty"`
	// 预存金淘宝商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolPreDepositGoldUnbindDto = sync.Pool{
	New: func() any {
		return new(PreDepositGoldUnbindDto)
	},
}

// GetPreDepositGoldUnbindDto() 从对象池中获取PreDepositGoldUnbindDto
func GetPreDepositGoldUnbindDto() *PreDepositGoldUnbindDto {
	return poolPreDepositGoldUnbindDto.Get().(*PreDepositGoldUnbindDto)
}

// ReleasePreDepositGoldUnbindDto 释放PreDepositGoldUnbindDto
func ReleasePreDepositGoldUnbindDto(v *PreDepositGoldUnbindDto) {
	v.OuterSalesActivityId = ""
	v.ItemId = 0
	poolPreDepositGoldUnbindDto.Put(v)
}
