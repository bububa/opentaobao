package wdk

// WdkOpenPurchasePrice 结构体
type WdkOpenPurchasePrice struct {
	// 经营店id，必填
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 淘系主订单号，必填
	TbOrderId string `json:"tb_order_id,omitempty" xml:"tb_order_id,omitempty"`
	// 子单信息
	WdkOpenPurchasePriceSubs []Wdkopenpurchasepricesubs `json:"wdk_open_purchase_price_subs,omitempty" xml:"wdk_open_purchase_price_subs>wdkopenpurchasepricesubs,omitempty"`
}
