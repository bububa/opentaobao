package alsc

// PropertyVerifyOpenReq 结构体
type PropertyVerifyOpenReq struct {
	// saas品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 操作人id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 外部id类型，wechat微信，alipay支付宝
	OuterType string `json:"outer_type,omitempty" xml:"outer_type,omitempty"`
	// 需要核销的积分
	Point int64 `json:"point,omitempty" xml:"point,omitempty"`
	// 请求幂等id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// saas门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 需要核销的储值，单位为分
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 券实例id
	VoucherIdList []string `json:"voucher_id_list,omitempty" xml:"voucher_id_list>string,omitempty"`
	// 需要核销的卡号id，不填默认为会员卡id
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 顾客id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
}
