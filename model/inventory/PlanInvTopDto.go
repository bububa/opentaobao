package inventory

// PlanInvTopDTO 
type PlanInvTopDTO struct {
    // 实例列表
    PlanInstanceInvList   []PlanInstanceInvResultDTO `json:"plan_instance_inv_list,omitempty" xml:"plan_instance_inv_list>plan_instance_inv_result_dto,omitempty"`
}
