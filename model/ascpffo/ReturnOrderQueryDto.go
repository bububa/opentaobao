package ascpffo

// ReturnOrderQueryDto 结构体
type ReturnOrderQueryDto struct {
	// 退供单号
	ReturnOrderNo string `json:"return_order_no,omitempty" xml:"return_order_no,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}
