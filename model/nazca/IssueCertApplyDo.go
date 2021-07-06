package nazca

// IssueCertApplyDo 结构体
type IssueCertApplyDo struct {
	// 合同编号
	ContractNum string `json:"contract_num,omitempty" xml:"contract_num,omitempty"`
	// 客户在1688的唯一标识
	PlatformUserId string `json:"platform_user_id,omitempty" xml:"platform_user_id,omitempty"`
	// 页面跳转同步通知URL
	ReturnUrl string `json:"return_url,omitempty" xml:"return_url,omitempty"`
	// 合同名称
	Topic string `json:"topic,omitempty" xml:"topic,omitempty"`
	// pageNum
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 角色 0：接收者 1：发送者
	SendReceive int64 `json:"send_receive,omitempty" xml:"send_receive,omitempty"`
}
