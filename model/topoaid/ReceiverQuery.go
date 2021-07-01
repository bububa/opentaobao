package topoaid

// ReceiverQuery 结构体
type ReceiverQuery struct {
	// 收件人ID (Open Addressee ID)，长度在128个字符之内。
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 交易订单ID
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
}
