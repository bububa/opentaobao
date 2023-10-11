package mtopopen

// ModifyRequest 结构体
type ModifyRequest struct {
	// 订单ID
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 快递公司标准编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 修改后收件人
	Fetcher string `json:"fetcher,omitempty" xml:"fetcher,omitempty"`
	// 修改后派送时间
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
}
