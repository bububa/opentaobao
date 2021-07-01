package rhino

// ClothingOutboundWaybill 结构体
type ClothingOutboundWaybill struct {
	// 每个运货单sku明细items
	OutboundItems []ClothingSkuDto `json:"outbound_items,omitempty" xml:"outbound_items>clothing_sku_dto,omitempty"`
	// 快递公司编码tms_code
	ExpressCompany string `json:"express_company,omitempty" xml:"express_company,omitempty"`
	// 快递编号tms_order_code
	ExpressId string `json:"express_id,omitempty" xml:"express_id,omitempty"`
}
