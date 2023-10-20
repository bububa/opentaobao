package wms

import (
	"sync"
)

// Consignsendinfo 结构体
type Consignsendinfo struct {
	// 发票确认信息列表
	InvoinceConfirmList []Invoinceconfirmlist `json:"invoince_confirm_list,omitempty" xml:"invoince_confirm_list>invoinceconfirmlist,omitempty"`
	// 运单信息
	TmsOrderList []Tmsorderlist `json:"tms_order_list,omitempty" xml:"tms_order_list>tmsorderlist,omitempty"`
	// 订单信息
	OrderItemList []Orderitemlist `json:"order_item_list,omitempty" xml:"order_item_list>orderitemlist,omitempty"`
	// 仓库订单完成时间
	ConfirmTime string `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
	// 订单状态 WMS_ACCEPT 接单成功 WMS_REJECT 接单失败 WMS_CONFIRMED 仓库生产完成
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 仓储编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 菜鸟订单编码
	CnOrderCode string `json:"cn_order_code,omitempty" xml:"cn_order_code,omitempty"`
	// ERP订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 订单类型 201 销售出库 502 换货出库 503 补发出库
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
}

var poolConsignsendinfo = sync.Pool{
	New: func() any {
		return new(Consignsendinfo)
	},
}

// GetConsignsendinfo() 从对象池中获取Consignsendinfo
func GetConsignsendinfo() *Consignsendinfo {
	return poolConsignsendinfo.Get().(*Consignsendinfo)
}

// ReleaseConsignsendinfo 释放Consignsendinfo
func ReleaseConsignsendinfo(v *Consignsendinfo) {
	v.InvoinceConfirmList = v.InvoinceConfirmList[:0]
	v.TmsOrderList = v.TmsOrderList[:0]
	v.OrderItemList = v.OrderItemList[:0]
	v.ConfirmTime = ""
	v.Status = ""
	v.StoreCode = ""
	v.CnOrderCode = ""
	v.OrderCode = ""
	v.OrderType = 0
	poolConsignsendinfo.Put(v)
}
