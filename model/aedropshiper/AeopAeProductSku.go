package aedropshiper

// AeopAeProductSku 结构体
type AeopAeProductSku struct {
	// SKU属性对象列表
	AeopSKUPropertyList []AeopSkuProperty `json:"aeop_s_k_u_property_list,omitempty" xml:"aeop_s_k_u_property_list>aeop_sku_property,omitempty"`
	// Sku属性对象list，允许1-3个sku属性对象，按sku属性顺序排放。sku属性从类目属性接口getAttributesResultByCateId获取。该项值输入sku属性，不能输入普通类目属性。注意，sku属性是有顺序的，必须按照顺序存放。
	AeopSKUPropertys []AeopSkuProperty `json:"aeop_s_k_u_propertys,omitempty" xml:"aeop_s_k_u_propertys>aeop_sku_property,omitempty"`
	// Sku价格。取值范围:0.01-100000;单位:美元。 如:200.07，表示:200美元7分。需要在正确的价格区间内。
	SkuPrice string `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
	// SKU ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 产品的货币单位。美元: USD, 卢布: RUB
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// sku折扣价格
	OfferSalePrice string `json:"offer_sale_price,omitempty" xml:"offer_sale_price,omitempty"`
	// sku批量折扣价
	OfferBulkSalePrice string `json:"offer_bulk_sale_price,omitempty" xml:"offer_bulk_sale_price,omitempty"`
	// Sku商家编码。 格式:半角英数字,长度20,不包含空格大于号和小于号。如果用户只填写零售价（productprice）和商品编码，需要完整生成一条SKU记录提交，否则商品编码无法保存。系统会认为只提交了零售价，而没有SKU，导致商品编辑未保存。
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// sku库存
	SKUAvailableStock int64 `json:"s_k_u_available_stock,omitempty" xml:"s_k_u_available_stock,omitempty"`
	// SKU实际可售库存属性ipmSkuStock，该属性值的合理取值范围为0~999999，如该商品有SKU时，请确保至少有一个SKU是有货状态，也就是ipmSkuStock取值是1~999999，在整个商品纬度库存值的取值范围是1~999999。 如果同时设置了skuStock属性，那么系统以ipmSkuStock属性为优先；如果没有设置ipmSkuStock属性，那么系统会根据skuStock属性进行设置库存，true表示999，false表示0。
	IpmSkuStock int64 `json:"ipm_sku_stock,omitempty" xml:"ipm_sku_stock,omitempty"`
	// 最少起批的数量
	SkuBulkOrder int64 `json:"sku_bulk_order,omitempty" xml:"sku_bulk_order,omitempty"`
	// Sku库存,数据格式有货true，无货false；至少有一条sku记录是有货的。
	SkuStock bool `json:"sku_stock,omitempty" xml:"sku_stock,omitempty"`
}
