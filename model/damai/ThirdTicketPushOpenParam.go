package damai

// ThirdTicketPushOpenParam 结构体
type ThirdTicketPushOpenParam struct {
	// 证件号
	CertificateNo string `json:"certificate_no,omitempty" xml:"certificate_no,omitempty"`
	// 用户手机号
	OrderUserMobile string `json:"order_user_mobile,omitempty" xml:"order_user_mobile,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 二维码
	QrCode string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 换票码
	TicketChangeCode string `json:"ticket_change_code,omitempty" xml:"ticket_change_code,omitempty"`
	// 用户名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 价格内容
	PriceContent string `json:"price_content,omitempty" xml:"price_content,omitempty"`
	// 票面元素扩展字段
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// rfid值
	Rfid string `json:"rfid,omitempty" xml:"rfid,omitempty"`
	// 证件内容
	CertificateType int64 `json:"certificate_type,omitempty" xml:"certificate_type,omitempty"`
	// 票面id
	FaceId int64 `json:"face_id,omitempty" xml:"face_id,omitempty"`
	// 楼层id
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 座位号
	SeatColId int64 `json:"seat_col_id,omitempty" xml:"seat_col_id,omitempty"`
	// 座位排号
	SeatRowId int64 `json:"seat_row_id,omitempty" xml:"seat_row_id,omitempty"`
	// 座位类型
	SeatType int64 `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// 看台id
	StandId int64 `json:"stand_id,omitempty" xml:"stand_id,omitempty"`
	// 商户id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 票品id
	TicketItemId int64 `json:"ticket_item_id,omitempty" xml:"ticket_item_id,omitempty"`
	// 票类型
	TicketType int64 `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
	// 票单号
	VoucherId int64 `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 票类型
	MediumType int64 `json:"medium_type,omitempty" xml:"medium_type,omitempty"`
	// 纸质票票单换票状态:1=未换，2=已换，
	PrintStatus int64 `json:"print_status,omitempty" xml:"print_status,omitempty"`
}
