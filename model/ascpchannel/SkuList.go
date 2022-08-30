package ascpchannel

// SkuList 结构体
type SkuList struct {
	// 经营模式
	SalesModeList []string `json:"sales_mode_list,omitempty" xml:"sales_mode_list>string,omitempty"`
	// 图片链接
	PictureList []string `json:"picture_list,omitempty" xml:"picture_list>string,omitempty"`
	// 销售属性
	Properties []string `json:"properties,omitempty" xml:"properties>string,omitempty"`
	// sku 对应的货品 id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// skuid
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
