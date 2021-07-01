package inventory

// PlanInvTopDto 
type PlanInvTopDto struct {
    // 实例列表
    PlanInstanceInvList   []PlanInstanceInvResultDto `json:"plan_instance_inv_list,omitempty" xml:"plan_instance_inv_list>plan_instance_inv_result_dto,omitempty"`
}
