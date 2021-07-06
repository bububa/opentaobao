package alicom

// FlowSuppllierNotifyModel 结构体
type FlowSuppllierNotifyModel struct {
	// 阿里通信业务id，具体询问技术接口人
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 用于区分商家通知的内容类型，例如：“0”为套餐激活通知
	NotifyType string `json:"notify_type,omitempty" xml:"notify_type,omitempty"`
	// 套餐生效时间
	OfferInvalidDate string `json:"offer_invalid_date,omitempty" xml:"offer_invalid_date,omitempty"`
	// 套餐失效时间
	OfferValidDate string `json:"offer_valid_date,omitempty" xml:"offer_valid_date,omitempty"`
	// 淘宝内部订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 运营商外部订单号
	OutOrderNo string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	// iccid和卡的过期时间
	SendGoodParam string `json:"send_good_param,omitempty" xml:"send_good_param,omitempty"`
	// iccid
	Iccid string `json:"iccid,omitempty" xml:"iccid,omitempty"`
	// 买家ID，用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 运营商的收货状态，false：未收到
	ReceiveStatus bool `json:"receive_status,omitempty" xml:"receive_status,omitempty"`
}
