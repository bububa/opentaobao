package idle

// AlibabaidlerentitemqueryData 结构体
type AlibabaidlerentitemqueryData struct {
	// 商品sku信息
	ItemSkuList []ItemSkuDto `json:"item_sku_list,omitempty" xml:"item_sku_list>item_sku_dto,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 新旧水平
	UsedLevel int64 `json:"used_level,omitempty" xml:"used_level,omitempty"`
	// 库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 地址信息
	Address *AddressDto `json:"address,omitempty" xml:"address,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 价格信息
	PriceInfo *PriceDto `json:"price_info,omitempty" xml:"price_info,omitempty"`
	// 商品状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 运费模板id
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}
