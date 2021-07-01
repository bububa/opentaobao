package trade

// OrderQuery 结构体
type OrderQuery struct {
	// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
	StartCreated string `json:"start_created,omitempty" xml:"start_created,omitempty"`
	// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
	EndCreated string `json:"end_created,omitempty" xml:"end_created,omitempty"`
	// 收件人的姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 收件人的手机号
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 收件人电话号码
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
}
