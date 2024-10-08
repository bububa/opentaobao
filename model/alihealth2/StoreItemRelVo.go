package alihealth2

import (
	"sync"
)

// StoreItemRelVo 结构体
type StoreItemRelVo struct {
	// 审核原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 审核状态信息
	CheckStatusString string `json:"check_status_string,omitempty" xml:"check_status_string,omitempty"`
	// ISV门店ID
	SpStoreId string `json:"sp_store_id,omitempty" xml:"sp_store_id,omitempty"`
	// ISV商品ID
	SpItemId string `json:"sp_item_id,omitempty" xml:"sp_item_id,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 绑定ID
	BindId int64 `json:"bind_id,omitempty" xml:"bind_id,omitempty"`
	// 审核状态码
	CheckStatus int64 `json:"check_status,omitempty" xml:"check_status,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolStoreItemRelVo = sync.Pool{
	New: func() any {
		return new(StoreItemRelVo)
	},
}

// GetStoreItemRelVo() 从对象池中获取StoreItemRelVo
func GetStoreItemRelVo() *StoreItemRelVo {
	return poolStoreItemRelVo.Get().(*StoreItemRelVo)
}

// ReleaseStoreItemRelVo 释放StoreItemRelVo
func ReleaseStoreItemRelVo(v *StoreItemRelVo) {
	v.Reason = ""
	v.CheckStatusString = ""
	v.SpStoreId = ""
	v.SpItemId = ""
	v.StoreId = 0
	v.BindId = 0
	v.CheckStatus = 0
	v.ItemId = 0
	poolStoreItemRelVo.Put(v)
}
