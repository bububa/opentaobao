package dengta

// ReturnMessage 结构体
type ReturnMessage struct {
	// 错误信息
	Cust string `json:"cust,omitempty" xml:"cust,omitempty"`
	// 错误信息
	Dev string `json:"dev,omitempty" xml:"dev,omitempty"`
}
