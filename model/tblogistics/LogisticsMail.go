package tblogistics

// LogisticsMail 结构体
type LogisticsMail struct {
	// 运单号.具体一个物流公司的运单号码.
	OutSid string `json:"out_sid,omitempty" xml:"out_sid,omitempty"`
	// 物流公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
}
