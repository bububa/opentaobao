package flight

// AuxRefundApiBean 结构体
type AuxRefundApiBean struct {
	// 是否可单独退 true可以，false不可以 当canRefund =true时此节点必传
	CanRefundIndependent bool `json:"can_refund_independent,omitempty" xml:"can_refund_independent,omitempty"`
	// 改规则。 最多允许200个字符。 禁止空格等特殊符号。
	ModifyRule string `json:"modify_rule,omitempty" xml:"modify_rule,omitempty"`
	// 是否可改 true可以，false不可以 当refundRule=true时此节点必传
	CanModify bool `json:"can_modify,omitempty" xml:"can_modify,omitempty"`
	// 是否可单独改 true可以，false不可以
	CanModifyIndependent bool `json:"can_modify_independent,omitempty" xml:"can_modify_independent,omitempty"`
	// 是否可退 true 可以，false不可以
	CanRefund bool `json:"can_refund,omitempty" xml:"can_refund,omitempty"`
	// 退规则。 最多允许200个字符。 禁止空格等特殊符号。
	RefundRule string `json:"refund_rule,omitempty" xml:"refund_rule,omitempty"`
}
