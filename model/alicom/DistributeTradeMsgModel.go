package alicom

// DistributeTradeMsgModel 结构体
type DistributeTradeMsgModel struct {
	// 业务来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 业务类型
	WtBizType string `json:"wt_biz_type,omitempty" xml:"wt_biz_type,omitempty"`
	// 要办理的账号
	Account string `json:"account,omitempty" xml:"account,omitempty"`
	// 扩展参数-短信验证码信息
	WttParam string `json:"wtt_param,omitempty" xml:"wtt_param,omitempty"`
	// 商家唯一幂等流水，数字类型，不可超过15位（建议appkey+数值）
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 商品sku，可空
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 要办理的商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
