package tmallhk

import (
	"sync"
)

// CtsInfo 结构体
type CtsInfo struct {
	// 货品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 商品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 子订单号
	SubOrderNo string `json:"sub_order_no,omitempty" xml:"sub_order_no,omitempty"`
	// 溯源码
	TraceCode string `json:"trace_code,omitempty" xml:"trace_code,omitempty"`
	// 托运
	Carriage *CtsCarriage `json:"carriage,omitempty" xml:"carriage,omitempty"`
	// 成品证书
	CompletedNgtc *CtsNgtc `json:"completed_ngtc,omitempty" xml:"completed_ngtc,omitempty"`
	// 国内物流
	Delivery *CtsDelivery `json:"delivery,omitempty" xml:"delivery,omitempty"`
	// 裸钻证书
	DiamondNgtc *CtsNgtc `json:"diamond_ngtc,omitempty" xml:"diamond_ngtc,omitempty"`
	// 戒托信息
	Ring *CtsRing `json:"ring,omitempty" xml:"ring,omitempty"`
	// 国内报关
	Shipment *CtsShipment `json:"shipment,omitempty" xml:"shipment,omitempty"`
}

var poolCtsInfo = sync.Pool{
	New: func() any {
		return new(CtsInfo)
	},
}

// GetCtsInfo() 从对象池中获取CtsInfo
func GetCtsInfo() *CtsInfo {
	return poolCtsInfo.Get().(*CtsInfo)
}

// ReleaseCtsInfo 释放CtsInfo
func ReleaseCtsInfo(v *CtsInfo) {
	v.ItemId = ""
	v.OrderNo = ""
	v.ProductId = ""
	v.SubOrderNo = ""
	v.TraceCode = ""
	v.Carriage = nil
	v.CompletedNgtc = nil
	v.Delivery = nil
	v.DiamondNgtc = nil
	v.Ring = nil
	v.Shipment = nil
	poolCtsInfo.Put(v)
}
