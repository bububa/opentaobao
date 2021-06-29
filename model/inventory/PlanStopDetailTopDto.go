package inventory

// PlanStopDetailTopDTO 
type PlanStopDetailTopDTO struct {
    // 外部商家系统单据号，用于生成计划库存的业务来源
    PlanOrderId   string `json:"plan_order_id,omitempty" xml:"plan_order_id,omitempty"`
    // 操作码
    OperateCode   string `json:"operate_code,omitempty" xml:"operate_code,omitempty"`
}
