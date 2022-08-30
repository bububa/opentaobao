package fenxiao

// RefundLogistics 结构体
type RefundLogistics struct {
	// 退货物流公司编码，如顺丰、韵达等
	CompanyCode string `json:"company_code,omitempty" xml:"company_code,omitempty"`
	// 退货物流公司名称，如顺丰
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 退货物流运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
}
