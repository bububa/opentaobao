package logistic

// Shipping 结构体
type Shipping struct {
	// 返回发货是否成功。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
