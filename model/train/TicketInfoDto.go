package train

// TicketInfoDto 结构体
type TicketInfoDto struct {
	// 乘车人姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 乘车人类型编码（1 成人 2 儿童 3 学生 4 残、军）
	PassengerTypeCode string `json:"passenger_type_code,omitempty" xml:"passenger_type_code,omitempty"`
	// 证件类型（1 身份证 C 港澳通行证 G 台胞证 B 护照）
	CertificateTypeCode string `json:"certificate_type_code,omitempty" xml:"certificate_type_code,omitempty"`
	// 证件号
	CertificateNo string `json:"certificate_no,omitempty" xml:"certificate_no,omitempty"`
	// 手机号
	MobileNo string `json:"mobile_no,omitempty" xml:"mobile_no,omitempty"`
	// ttp子单号
	TtpSubOrderId int64 `json:"ttp_sub_order_id,omitempty" xml:"ttp_sub_order_id,omitempty"`
}
