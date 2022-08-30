package alscmerchant

// ExternalTicketUseRequest 结构体
type ExternalTicketUseRequest struct {
	// 特殊可选 核销份数，次卡业务必填、非次卡业务选填， 非次卡场景表示核销同一订单下的同类凭证的个数
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 特殊可选 口碑订单号，使用外部凭证码核销时必填
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 特殊可选 凭证码类型,INTERNAL_CODE(内部凭证码),EXTERNAL_CODE(外部凭证码),默认INTERNAL_CODE,外部券场景必填
	CodeType string `json:"code_type,omitempty" xml:"code_type,omitempty"`
	// 外部请求号，支持英文字母和数字，由开发者自行定义（不允许重复）
	TicketRequestId string `json:"ticket_request_id,omitempty" xml:"ticket_request_id,omitempty"`
	// 核销的口碑门店id(门店同步API中的alsc_store_id)，目前是必填。只有加白的合作方法可以不填写
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 门店id类型,INTERNAL_SHOP(内部店铺id),EXTERNAL_SHOP(外部店铺id),默认INTERNAL_SHOP
	ShopType string `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
	// 凭证码，包括内部凭证码和外部凭证码，内部凭证码为12位，纯数字，且唯一不重复
	TicketCode string `json:"ticket_code,omitempty" xml:"ticket_code,omitempty"`
	// 业务发生时间，默认为接口调用时间
	GmtBiz string `json:"gmt_biz,omitempty" xml:"gmt_biz,omitempty"`
	// 本地侧凭证id，如果是三方码值核销传参ticketId,则该参数必须
	TicketId string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
}
