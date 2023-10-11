package tuanhotel

// TuanItemOnlineSkuManagerVo 结构体
type TuanItemOnlineSkuManagerVo struct {
	// 套餐名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// sku id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 间夜
	Nights int64 `json:"nights,omitempty" xml:"nights,omitempty"`
}
