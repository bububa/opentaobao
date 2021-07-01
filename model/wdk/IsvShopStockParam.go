package wdk

// IsvShopStockParam 结构体
type IsvShopStockParam struct {
	// 商品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 门店编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 派样活动id
	SampleActivityId int64 `json:"sample_activity_id,omitempty" xml:"sample_activity_id,omitempty"`
}
