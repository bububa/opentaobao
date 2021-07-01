package nlife

// RetailSkuTopDo 结构体
type RetailSkuTopDo struct {
	// 上次修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// created
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 淘宝上的skuId
	TaobaoSkuId string `json:"taobao_sku_id,omitempty" xml:"taobao_sku_id,omitempty"`
	// 淘宝上的itemId
	TaobaoItemId string `json:"taobao_item_id,omitempty" xml:"taobao_item_id,omitempty"`
	// sku主图URL
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// sku价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// sku库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// sku属性名称
	PropertiesName string `json:"properties_name,omitempty" xml:"properties_name,omitempty"`
	// sku属性
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// barcode
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 外部商家的编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// sku的挂牌价-单位元,保留2位小数
	TagPrice string `json:"tag_price,omitempty" xml:"tag_price,omitempty"`
	// skuRefId
	SkuRefId int64 `json:"sku_ref_id,omitempty" xml:"sku_ref_id,omitempty"`
	// 网直供库存
	OnlineQuantity int64 `json:"online_quantity,omitempty" xml:"online_quantity,omitempty"`
	// 自用编码
	CustomCode string `json:"custom_code,omitempty" xml:"custom_code,omitempty"`
}
