package wdk

// PreCheckRequest 结构体
type PreCheckRequest struct {
	// 购买商品列表
	SkuList []PromiseSkuInfo `json:"sku_list,omitempty" xml:"sku_list>promise_sku_info,omitempty"`
	// 渠道店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 配送地址经纬度
	Geo string `json:"geo,omitempty" xml:"geo,omitempty"`
	// CartScene = 购物车场景 / PostOrderScene = 进入支付场景
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 选择时间片日期, PostOrderScene场景必填
	TimeSliceDate string `json:"time_slice_date,omitempty" xml:"time_slice_date,omitempty"`
	// 选择时间片时间段, PostOrderScene场景必填
	TimeSlice string `json:"time_slice,omitempty" xml:"time_slice,omitempty"`
}
