package flight

// IssueProOrderVo 结构体
type IssueProOrderVo struct {
	// 辅营出票号，不可为空，字符长度不超过32位
	ExtraNo string `json:"extra_no,omitempty" xml:"extra_no,omitempty"`
	// 辅营订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
