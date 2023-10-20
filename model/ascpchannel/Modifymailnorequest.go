package ascpchannel

import (
	"sync"
)

// Modifymailnorequest 结构体
type Modifymailnorequest struct {
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 原运单号
	OldTmsOrderCode string `json:"old_tms_order_code,omitempty" xml:"old_tms_order_code,omitempty"`
	// 新运单号
	NewTmsOrderCode string `json:"new_tms_order_code,omitempty" xml:"new_tms_order_code,omitempty"`
	// 原配送公司编码
	OldTmsServiceCode string `json:"old_tms_service_code,omitempty" xml:"old_tms_service_code,omitempty"`
	// 发货LP
	BizOrderCode string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
}

var poolModifymailnorequest = sync.Pool{
	New: func() any {
		return new(Modifymailnorequest)
	},
}

// GetModifymailnorequest() 从对象池中获取Modifymailnorequest
func GetModifymailnorequest() *Modifymailnorequest {
	return poolModifymailnorequest.Get().(*Modifymailnorequest)
}

// ReleaseModifymailnorequest 释放Modifymailnorequest
func ReleaseModifymailnorequest(v *Modifymailnorequest) {
	v.SupplierId = ""
	v.OldTmsOrderCode = ""
	v.NewTmsOrderCode = ""
	v.OldTmsServiceCode = ""
	v.BizOrderCode = ""
	poolModifymailnorequest.Put(v)
}
