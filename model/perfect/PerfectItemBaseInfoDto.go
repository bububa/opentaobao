package perfect

// PerfectItemBaseInfoDto 结构体
type PerfectItemBaseInfoDto struct {
	// 商品条码
	ItemBarcode string `json:"item_barcode,omitempty" xml:"item_barcode,omitempty"`
	// 商品外部编码
	ItemOuterId string `json:"item_outer_id,omitempty" xml:"item_outer_id,omitempty"`
	// 商品吊牌价
	ItemPretium string `json:"item_pretium,omitempty" xml:"item_pretium,omitempty"`
	// 商品价格
	ItemPrice string `json:"item_price,omitempty" xml:"item_price,omitempty"`
	// 商品体积
	ItemSize string `json:"item_size,omitempty" xml:"item_size,omitempty"`
	// 商品重量
	ItemWeight string `json:"item_weight,omitempty" xml:"item_weight,omitempty"`
	// 销售渠道
	SaleChannelKey string `json:"sale_channel_key,omitempty" xml:"sale_channel_key,omitempty"`
	// 商品数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
}
