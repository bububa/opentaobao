package wdk

// TimeSliceGetRequest 结构体
type TimeSliceGetRequest struct {
	// 购买商品列表
	SkuList []PromiseSkuInfo `json:"sku_list,omitempty" xml:"sku_list>promise_sku_info,omitempty"`
	// 渠道店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 配送地址经纬度
	Geo string `json:"geo,omitempty" xml:"geo,omitempty"`
}
