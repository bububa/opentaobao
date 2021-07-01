package miniappopen

// DistributionOrderBindTargetEntityOpenRequestV2 结构体
type DistributionOrderBindTargetEntityOpenRequestV2 struct {
	// 绑定/解绑的投放计划id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// true:绑定； false: 解绑
	AddBind bool `json:"add_bind,omitempty" xml:"add_bind,omitempty"`
}
