package moscm

import (
	"sync"
)

// SalesOrderItemDto 结构体
type SalesOrderItemDto struct {
	// 外部商品Id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 商品Id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品类型:普通商品(NORMAL)
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 销售属性：颜色:15#/规格:3.5g，用/分隔
	SaleProperty string `json:"sale_property,omitempty" xml:"sale_property,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 价格，单位:元
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 折扣金额，单位:元
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 吊牌价，单位:元
	TagPrice string `json:"tag_price,omitempty" xml:"tag_price,omitempty"`
	// 支付信息
	Payments string `json:"payments,omitempty" xml:"payments,omitempty"`
	// 唯一码（仅西有用）
	DiscCode string `json:"disc_code,omitempty" xml:"disc_code,omitempty"`
	// 分摊金额（仅西有用）
	DiscAmount string `json:"disc_amount,omitempty" xml:"disc_amount,omitempty"`
	// 发货拦截标识（ NORMAL 正常 ;INTERCEPTION 拦截 ）
	ShipmentInterception string `json:"shipment_interception,omitempty" xml:"shipment_interception,omitempty"`
	// 导购Id
	GuiderId string `json:"guider_id,omitempty" xml:"guider_id,omitempty"`
	// 条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 货号
	ArtNo string `json:"art_no,omitempty" xml:"art_no,omitempty"`
	// 唯一标识
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// E2分摊金额
	E2PartDiscount string `json:"e2_part_discount,omitempty" xml:"e2_part_discount,omitempty"`
	// 商品编码
	SettlementCode string `json:"settlement_code,omitempty" xml:"settlement_code,omitempty"`
}

var poolSalesOrderItemDto = sync.Pool{
	New: func() any {
		return new(SalesOrderItemDto)
	},
}

// GetSalesOrderItemDto() 从对象池中获取SalesOrderItemDto
func GetSalesOrderItemDto() *SalesOrderItemDto {
	return poolSalesOrderItemDto.Get().(*SalesOrderItemDto)
}

// ReleaseSalesOrderItemDto 释放SalesOrderItemDto
func ReleaseSalesOrderItemDto(v *SalesOrderItemDto) {
	v.OutId = ""
	v.SkuId = ""
	v.Title = ""
	v.ProductType = ""
	v.SaleProperty = ""
	v.Quantity = ""
	v.Price = ""
	v.Discount = ""
	v.TagPrice = ""
	v.Payments = ""
	v.DiscCode = ""
	v.DiscAmount = ""
	v.ShipmentInterception = ""
	v.GuiderId = ""
	v.Barcode = ""
	v.ArtNo = ""
	v.Id = ""
	v.E2PartDiscount = ""
	v.SettlementCode = ""
	poolSalesOrderItemDto.Put(v)
}
