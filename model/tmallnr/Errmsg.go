package tmallnr

// Errmsg 结构体
type Errmsg struct {
	// 订单号
	Key int64 `json:"key,omitempty" xml:"key,omitempty"`
	// 错误编码
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
