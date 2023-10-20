package ascpchannel

import (
	"sync"
)

// Instorageresultrequest 结构体
type Instorageresultrequest struct {
	// 逆向物流单号
	LgOrderCode string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 入库时间
	OrderConfirmTime string `json:"order_confirm_time,omitempty" xml:"order_confirm_time,omitempty"`
	// 操作信息
	OperateInfo *Operateinfo `json:"operate_info,omitempty" xml:"operate_info,omitempty"`
}

var poolInstorageresultrequest = sync.Pool{
	New: func() any {
		return new(Instorageresultrequest)
	},
}

// GetInstorageresultrequest() 从对象池中获取Instorageresultrequest
func GetInstorageresultrequest() *Instorageresultrequest {
	return poolInstorageresultrequest.Get().(*Instorageresultrequest)
}

// ReleaseInstorageresultrequest 释放Instorageresultrequest
func ReleaseInstorageresultrequest(v *Instorageresultrequest) {
	v.LgOrderCode = ""
	v.SupplierId = ""
	v.OrderConfirmTime = ""
	v.OperateInfo = nil
	poolInstorageresultrequest.Put(v)
}
