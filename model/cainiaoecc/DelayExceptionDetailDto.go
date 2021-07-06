package cainiaoecc

// DelayExceptionDetailDto 结构体
type DelayExceptionDetailDto struct {
	// CP回复列表
	CpReplyList []string `json:"cp_reply_list,omitempty" xml:"cp_reply_list>string,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 异常类型code
	ExceptionCode string `json:"exception_code,omitempty" xml:"exception_code,omitempty"`
	// 异常类型名称
	ExceptionName string `json:"exception_name,omitempty" xml:"exception_name,omitempty"`
	// 商家Id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
