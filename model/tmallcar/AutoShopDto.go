package tmallcar

// AutoShopDto 结构体
type AutoShopDto struct {
	// 门店自用编码
	ShopOuterNum string `json:"shop_outer_num,omitempty" xml:"shop_outer_num,omitempty"`
	// 门店名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 一级经销商名称
	FirstLevelShopName string `json:"first_level_shop_name,omitempty" xml:"first_level_shop_name,omitempty"`
	// 门店所在省份
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 门店所在城市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 经销商开票代码
	ShopInvoiceCode string `json:"shop_invoice_code,omitempty" xml:"shop_invoice_code,omitempty"`
	// 一级经销商编码
	FirstLevelShopCode string `json:"first_level_shop_code,omitempty" xml:"first_level_shop_code,omitempty"`
	// 门店id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}
