package omniorder

import (
	"sync"
)

// SubOrder 结构体
type SubOrder struct {
	// 0表示无系统异常
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 异常描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 店铺Id, 可能是门店或者电商仓
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 店铺类型, 门店或者电商仓
	StoreType string `json:"store_type,omitempty" xml:"store_type,omitempty"`
	// 店铺名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 操作者
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 子订单Id
	SubOid int64 `json:"sub_oid,omitempty" xml:"sub_oid,omitempty"`
	// 主订单Id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolSubOrder = sync.Pool{
	New: func() any {
		return new(SubOrder)
	},
}

// GetSubOrder() 从对象池中获取SubOrder
func GetSubOrder() *SubOrder {
	return poolSubOrder.Get().(*SubOrder)
}

// ReleaseSubOrder 释放SubOrder
func ReleaseSubOrder(v *SubOrder) {
	v.Code = ""
	v.Message = ""
	v.StoreId = ""
	v.StoreType = ""
	v.StoreName = ""
	v.Operator = ""
	v.Attributes = ""
	v.SubOid = 0
	v.Tid = 0
	poolSubOrder.Put(v)
}
