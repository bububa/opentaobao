package damai

import (
	"sync"
)

// TicketItemIdOpenParam 结构体
type TicketItemIdOpenParam struct {
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 票品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}

var poolTicketItemIdOpenParam = sync.Pool{
	New: func() any {
		return new(TicketItemIdOpenParam)
	},
}

// GetTicketItemIdOpenParam() 从对象池中获取TicketItemIdOpenParam
func GetTicketItemIdOpenParam() *TicketItemIdOpenParam {
	return poolTicketItemIdOpenParam.Get().(*TicketItemIdOpenParam)
}

// ReleaseTicketItemIdOpenParam 释放TicketItemIdOpenParam
func ReleaseTicketItemIdOpenParam(v *TicketItemIdOpenParam) {
	v.SupplierSecret = ""
	v.ItemId = 0
	v.PerformId = 0
	v.SystemId = 0
	poolTicketItemIdOpenParam.Put(v)
}
