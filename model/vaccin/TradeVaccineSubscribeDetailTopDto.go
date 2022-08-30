package vaccin

// TradeVaccineSubscribeDetailTopDto 结构体
type TradeVaccineSubscribeDetailTopDto struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 业务订单ID
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 预约者姓名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 预约者手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 身份证号
	IdentityNo string `json:"identity_no,omitempty" xml:"identity_no,omitempty"`
	// 次要证件类型： 2护照 3港澳通行证
	SecondaryCardType string `json:"secondary_card_type,omitempty" xml:"secondary_card_type,omitempty"`
	// 次要证件号码
	SecondaryCardId string `json:"secondary_card_id,omitempty" xml:"secondary_card_id,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 关联订单购买的具体的SKU信息
	Sku string `json:"sku,omitempty" xml:"sku,omitempty"`
	// pov名称
	PovName string `json:"pov_name,omitempty" xml:"pov_name,omitempty"`
	// pov地址
	PovAddress string `json:"pov_address,omitempty" xml:"pov_address,omitempty"`
	// APP渠道
	AppChannel string `json:"app_channel,omitempty" xml:"app_channel,omitempty"`
	// 预约开始时间
	SubscribeStartTime string `json:"subscribe_start_time,omitempty" xml:"subscribe_start_time,omitempty"`
	// 预约结束时间
	SubscribeEndTime string `json:"subscribe_end_time,omitempty" xml:"subscribe_end_time,omitempty"`
	// 订单实付金额
	ActualPayFee string `json:"actual_pay_fee,omitempty" xml:"actual_pay_fee,omitempty"`
	// 取消者角色
	CancelRole string `json:"cancel_role,omitempty" xml:"cancel_role,omitempty"`
	// 预约单主键
	SubscribeId int64 `json:"subscribe_id,omitempty" xml:"subscribe_id,omitempty"`
	// 预约时段 1上午，2下午
	PeriodTag int64 `json:"period_tag,omitempty" xml:"period_tag,omitempty"`
	// 预约单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
