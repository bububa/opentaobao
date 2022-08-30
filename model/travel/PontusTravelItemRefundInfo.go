package travel

// PontusTravelItemRefundInfo 结构体
type PontusTravelItemRefundInfo struct {
	// 退改规则 1）格式：标准规则 或 自定义规则：a_a_num  2）规则：自定义规则里最多可含5组规则
	RefundRegulations []string `json:"refund_regulations,omitempty" xml:"refund_regulations>string,omitempty"`
	// 退改规则类型，0为标准，1为自定义，2为不支持退改规则。不传默认为0
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}
