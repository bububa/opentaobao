package alscmerchant

// TicketConsultResponse 结构体
type TicketConsultResponse struct {
	// 券码所属订单下，可使用的凭证（同商品）列表
	TicketInfoList []TicketInfo `json:"ticket_info_list,omitempty" xml:"ticket_info_list>ticket_info,omitempty"`
	// 脱敏手机号，例：*******4139。如果商品类目未开放电话组件，不返回数据
	UserPhone string `json:"user_phone,omitempty" xml:"user_phone,omitempty"`
	// 混淆后的用户uid。查询订单接口用
	BuyerId string `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 本地订单-主单号。长度：19位
	AlscOrderNo string `json:"alsc_order_no,omitempty" xml:"alsc_order_no,omitempty"`
	// 本地订单-子单号。长度：19位
	AlscSubOrderNo string `json:"alsc_sub_order_no,omitempty" xml:"alsc_sub_order_no,omitempty"`
	// 口碑商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 外部商品id
	OutItemId string `json:"out_item_id,omitempty" xml:"out_item_id,omitempty"`
	// 剩余可核销数量。团购举例：购买2份团购券，已核销了1份，该属性返回1；次卡举例：购买了3次卡，已核销1次，该属性返回2
	AvailableQuantity int64 `json:"available_quantity,omitempty" xml:"available_quantity,omitempty"`
}
