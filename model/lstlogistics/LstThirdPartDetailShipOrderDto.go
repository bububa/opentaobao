package lstlogistics

import (
	"sync"
)

// LstThirdPartDetailShipOrderDto 结构体
type LstThirdPartDetailShipOrderDto struct {
	// 销售单位
	SkuUnit string `json:"sku_unit,omitempty" xml:"sku_unit,omitempty"`
	// 规格
	SkuSpec string `json:"sku_spec,omitempty" xml:"sku_spec,omitempty"`
	// 条码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 明细订单ID
	DetailOrderId string `json:"detail_order_id,omitempty" xml:"detail_order_id,omitempty"`
	// 子发货单状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 销售数量
	SaleQuantity int64 `json:"sale_quantity,omitempty" xml:"sale_quantity,omitempty"`
	// 装车数量
	LoadQuantity int64 `json:"load_quantity,omitempty" xml:"load_quantity,omitempty"`
	// 签收数量
	SignQuantity int64 `json:"sign_quantity,omitempty" xml:"sign_quantity,omitempty"`
}

var poolLstThirdPartDetailShipOrderDto = sync.Pool{
	New: func() any {
		return new(LstThirdPartDetailShipOrderDto)
	},
}

// GetLstThirdPartDetailShipOrderDto() 从对象池中获取LstThirdPartDetailShipOrderDto
func GetLstThirdPartDetailShipOrderDto() *LstThirdPartDetailShipOrderDto {
	return poolLstThirdPartDetailShipOrderDto.Get().(*LstThirdPartDetailShipOrderDto)
}

// ReleaseLstThirdPartDetailShipOrderDto 释放LstThirdPartDetailShipOrderDto
func ReleaseLstThirdPartDetailShipOrderDto(v *LstThirdPartDetailShipOrderDto) {
	v.SkuUnit = ""
	v.SkuSpec = ""
	v.SkuCode = ""
	v.SkuName = ""
	v.DetailOrderId = ""
	v.Status = ""
	v.SaleQuantity = 0
	v.LoadQuantity = 0
	v.SignQuantity = 0
	poolLstThirdPartDetailShipOrderDto.Put(v)
}
