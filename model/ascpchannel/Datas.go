package ascpchannel

import (
	"sync"
)

// Datas 结构体
type Datas struct {
	// 占用库存
	LockQuantity string `json:"lock_quantity,omitempty" xml:"lock_quantity,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 总库存
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 供应链中台货仓code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 1-良品 101 -在库残次
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 供应链中台货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 供应链中台货主ID
	SourceUserId int64 `json:"source_user_id,omitempty" xml:"source_user_id,omitempty"`
}

var poolDatas = sync.Pool{
	New: func() any {
		return new(Datas)
	},
}

// GetDatas() 从对象池中获取Datas
func GetDatas() *Datas {
	return poolDatas.Get().(*Datas)
}

// ReleaseDatas 释放Datas
func ReleaseDatas(v *Datas) {
	v.LockQuantity = ""
	v.ChannelCode = ""
	v.Quantity = ""
	v.StoreCode = ""
	v.InventoryType = 0
	v.ScItemId = 0
	v.SourceUserId = 0
	poolDatas.Put(v)
}
