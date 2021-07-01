package wdk

// OrderCommitReq 结构体
type OrderCommitReq struct {
	// 商品列表
	ItemConfirmInfos []ItemConfirmInfo `json:"item_confirm_infos,omitempty" xml:"item_confirm_infos>item_confirm_info,omitempty"`
	// 订单号
	ExternalOrderNo string `json:"external_order_no,omitempty" xml:"external_order_no,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
}
