package rhino

import (
	"sync"
)

// ClothingOutboundWaybill 结构体
type ClothingOutboundWaybill struct {
	// 每个运货单sku明细items
	OutboundItems []ClothingSkuDto `json:"outbound_items,omitempty" xml:"outbound_items>clothing_sku_dto,omitempty"`
	// 快递公司编码tms_code
	ExpressCompany string `json:"express_company,omitempty" xml:"express_company,omitempty"`
	// 快递编号tms_order_code
	ExpressId string `json:"express_id,omitempty" xml:"express_id,omitempty"`
}

var poolClothingOutboundWaybill = sync.Pool{
	New: func() any {
		return new(ClothingOutboundWaybill)
	},
}

// GetClothingOutboundWaybill() 从对象池中获取ClothingOutboundWaybill
func GetClothingOutboundWaybill() *ClothingOutboundWaybill {
	return poolClothingOutboundWaybill.Get().(*ClothingOutboundWaybill)
}

// ReleaseClothingOutboundWaybill 释放ClothingOutboundWaybill
func ReleaseClothingOutboundWaybill(v *ClothingOutboundWaybill) {
	v.OutboundItems = v.OutboundItems[:0]
	v.ExpressCompany = ""
	v.ExpressId = ""
	poolClothingOutboundWaybill.Put(v)
}
