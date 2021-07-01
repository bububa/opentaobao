package bus

// B2BFetchHolderInfo 结构体
type B2BFetchHolderInfo struct {
	// 取票人证件号码
	FetchCertNumber string `json:"fetch_cert_number,omitempty" xml:"fetch_cert_number,omitempty"`
	// 取票人证件类型
	FetchCertType string `json:"fetch_cert_type,omitempty" xml:"fetch_cert_type,omitempty"`
	// 取票人电话
	FetchPhone string `json:"fetch_phone,omitempty" xml:"fetch_phone,omitempty"`
	// 取票人姓名
	FetchTicketName string `json:"fetch_ticket_name,omitempty" xml:"fetch_ticket_name,omitempty"`
}
