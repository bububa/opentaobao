package icbu

// CustomContent 结构体
type CustomContent struct {
	// 定制类型，只允许填写英文字符
	CustomType string `json:"custom_type,omitempty" xml:"custom_type,omitempty"`
	// 最小起订量
	MinOrderQuantity int64 `json:"min_order_quantity,omitempty" xml:"min_order_quantity,omitempty"`
}
