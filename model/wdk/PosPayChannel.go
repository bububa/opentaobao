package wdk

// PosPayChannel 结构体
type PosPayChannel struct {
	// 该支付方式对应的支付金额
	PayAmount int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// 支付方式编码，盒马给出了常见支付方式的编码
	PayType string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
}
