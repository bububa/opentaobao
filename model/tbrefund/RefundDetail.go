package tbrefund

// RefundDetail 结构体
type RefundDetail struct {
	// 退款当前可以执行的操作
	AllowedOperations []Operation `json:"allowed_operations,omitempty" xml:"allowed_operations>operation,omitempty"`
	// 退款当前不可以执行的操作
	NotAllowedOperations []Operation `json:"not_allowed_operations,omitempty" xml:"not_allowed_operations>operation,omitempty"`
	// 退款版本号
	RefundVersion int64 `json:"refund_version,omitempty" xml:"refund_version,omitempty"`
}
