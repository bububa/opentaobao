package inventory

// InvUnifyPlanTopQuerys 结构体
type InvUnifyPlanTopQuerys struct {
	// 查询入参
	PlanQueryList []InvUnifyPlanTopQuery `json:"plan_query_list,omitempty" xml:"plan_query_list>inv_unify_plan_top_query,omitempty"`
}
