package miniappopen

// DistributionOrderBindBaseDto 结构体
type DistributionOrderBindBaseDto struct {
	// 商品id
	TargetEntityId string `json:"target_entity_id,omitempty" xml:"target_entity_id,omitempty"`
	// 失败的原因
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 绑定是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
