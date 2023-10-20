package ascpchannel

import (
	"sync"
)

// Consignordershiprequest 结构体
type Consignordershiprequest struct {
	// 履约子单明细
	OrderItems []Orderitems `json:"order_items,omitempty" xml:"order_items>orderitems,omitempty"`
	// 包裹列表
	TmsOrders []Tmsorders `json:"tms_orders,omitempty" xml:"tms_orders>tmsorders,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 外部业务号，幂等控制使用
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 履约单号
	BizOrderCode string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
	// 发货仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 发货仓名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 一盘货业务模式，默认为0代表商家仓商家配，为1代表商家仓自营配 (为1时会强制校验配CP和单号必须与取号时一致，且多包裹必须一次性发货)
	BusinessModel string `json:"business_model,omitempty" xml:"business_model,omitempty"`
	// 发件人信息
	SenderInfo *Senderinfo `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
	// 是否整单发货,目前只支持履约单整单发货回传
	WholeSheetConsigned bool `json:"whole_sheet_consigned,omitempty" xml:"whole_sheet_consigned,omitempty"`
}

var poolConsignordershiprequest = sync.Pool{
	New: func() any {
		return new(Consignordershiprequest)
	},
}

// GetConsignordershiprequest() 从对象池中获取Consignordershiprequest
func GetConsignordershiprequest() *Consignordershiprequest {
	return poolConsignordershiprequest.Get().(*Consignordershiprequest)
}

// ReleaseConsignordershiprequest 释放Consignordershiprequest
func ReleaseConsignordershiprequest(v *Consignordershiprequest) {
	v.OrderItems = v.OrderItems[:0]
	v.TmsOrders = v.TmsOrders[:0]
	v.SupplierId = ""
	v.OutBizId = ""
	v.BizOrderCode = ""
	v.StoreCode = ""
	v.StoreName = ""
	v.BusinessModel = ""
	v.SenderInfo = nil
	v.WholeSheetConsigned = false
	poolConsignordershiprequest.Put(v)
}
