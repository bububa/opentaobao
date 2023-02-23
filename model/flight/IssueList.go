package flight

// IssueList 结构体
type IssueList struct {
	// 票号
	Tickets []string `json:"tickets,omitempty" xml:"tickets>string,omitempty"`
	// 税项对象
	Taxes []Taxes `json:"taxes,omitempty" xml:"taxes>taxes,omitempty"`
	// 航段
	Segments []Segments `json:"segments,omitempty" xml:"segments>segments,omitempty"`
	// 政策信息
	SellPolicyList []SellPolicyDto `json:"sell_policy_list,omitempty" xml:"sell_policy_list>sell_policy_dto,omitempty"`
	// 证件信息
	CertNo string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 乘机人姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// pnr
	Pnr string `json:"pnr,omitempty" xml:"pnr,omitempty"`
	// 联系电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 乘机人生日
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 大编码
	BigPnr string `json:"big_pnr,omitempty" xml:"big_pnr,omitempty"`
	// 证件类型
	CertType int64 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// 乘客类型
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 票面价
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 优惠价格
	Promotion int64 `json:"promotion,omitempty" xml:"promotion,omitempty"`
}
