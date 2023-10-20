package damai

import (
	"sync"
)

// TicketIdOpenParam 结构体
type TicketIdOpenParam struct {
	// 商家密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 票单id
	VoucherId int64 `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
}

var poolTicketIdOpenParam = sync.Pool{
	New: func() any {
		return new(TicketIdOpenParam)
	},
}

// GetTicketIdOpenParam() 从对象池中获取TicketIdOpenParam
func GetTicketIdOpenParam() *TicketIdOpenParam {
	return poolTicketIdOpenParam.Get().(*TicketIdOpenParam)
}

// ReleaseTicketIdOpenParam 释放TicketIdOpenParam
func ReleaseTicketIdOpenParam(v *TicketIdOpenParam) {
	v.SupplierSecret = ""
	v.PerformId = 0
	v.SystemId = 0
	v.VoucherId = 0
	poolTicketIdOpenParam.Put(v)
}
