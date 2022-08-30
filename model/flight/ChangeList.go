package flight

// ChangeList 结构体
type ChangeList struct {
	// 票号
	Tickets []string `json:"tickets,omitempty" xml:"tickets>string,omitempty"`
	// 改签前数据
	BeforeChangeSegments []BeforeChangeSegments `json:"before_change_segments,omitempty" xml:"before_change_segments>before_change_segments,omitempty"`
	// 改签后数据
	AfterChangeSegments []AfterChangeSegments `json:"after_change_segments,omitempty" xml:"after_change_segments>after_change_segments,omitempty"`
	// 改签前票号
	BeforeChangeTickets []string `json:"before_change_tickets,omitempty" xml:"before_change_tickets>string,omitempty"`
	// 证件号
	CertNo string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 改签后pnr
	Pnr string `json:"pnr,omitempty" xml:"pnr,omitempty"`
	// 证件类型:0:身份证, 1:护照, 2:学生证, 3:军官证, 4:回乡证, 5:台胞证, 6:港澳通行证, 7:国际海员, 8:外国人永久居留证, 9:其他证件, 10:警官证, 11:士兵证件, 12:台湾通行证, 13:入台证, 14:户口簿, 15:出生证, 16:驾驶证, 17:港澳居民居住证, 18:台湾居民居住证,
	CertType int64 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// 乘客类型: 1:成人, 2:儿童, 3:婴儿, 4:留学生
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 优惠;单位:分
	Promotion int64 `json:"promotion,omitempty" xml:"promotion,omitempty"`
	// 票面价;单位:分
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 改签费;单位:分
	ChangeFee int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// 升舱费;单位:分
	UpgradeFee int64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}
