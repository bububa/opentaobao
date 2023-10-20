package lstlogistics

import (
	"sync"
)

// LstThirdPartDetailShipOrderCreateDto 结构体
type LstThirdPartDetailShipOrderCreateDto struct {
	// 品牌
	SkuBrand string `json:"sku_brand,omitempty" xml:"sku_brand,omitempty"`
	// 规格
	SkuSpec string `json:"sku_spec,omitempty" xml:"sku_spec,omitempty"`
	// 销售单位
	SkuUnit string `json:"sku_unit,omitempty" xml:"sku_unit,omitempty"`
	// 国标条码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 订单明细id，可用来标注唯一的订单明细行。当其他无法唯一标识的时候，可用此字段。
	DetailOrderId string `json:"detail_order_id,omitempty" xml:"detail_order_id,omitempty"`
	// 购买数量
	SaleQuantity int64 `json:"sale_quantity,omitempty" xml:"sale_quantity,omitempty"`
	// 单价，单位为分
	SkuPrice int64 `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
	// 货品实付金额，单位为分
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
}

var poolLstThirdPartDetailShipOrderCreateDto = sync.Pool{
	New: func() any {
		return new(LstThirdPartDetailShipOrderCreateDto)
	},
}

// GetLstThirdPartDetailShipOrderCreateDto() 从对象池中获取LstThirdPartDetailShipOrderCreateDto
func GetLstThirdPartDetailShipOrderCreateDto() *LstThirdPartDetailShipOrderCreateDto {
	return poolLstThirdPartDetailShipOrderCreateDto.Get().(*LstThirdPartDetailShipOrderCreateDto)
}

// ReleaseLstThirdPartDetailShipOrderCreateDto 释放LstThirdPartDetailShipOrderCreateDto
func ReleaseLstThirdPartDetailShipOrderCreateDto(v *LstThirdPartDetailShipOrderCreateDto) {
	v.SkuBrand = ""
	v.SkuSpec = ""
	v.SkuUnit = ""
	v.SkuCode = ""
	v.SkuName = ""
	v.DetailOrderId = ""
	v.SaleQuantity = 0
	v.SkuPrice = 0
	v.PayFee = 0
	poolLstThirdPartDetailShipOrderCreateDto.Put(v)
}
