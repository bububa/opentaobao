package inventory

import (
	"sync"
)

// InvUnifyPlanTopQuerys 结构体
type InvUnifyPlanTopQuerys struct {
	// 查询入参
	PlanQueryList []InvUnifyPlanTopQuery `json:"plan_query_list,omitempty" xml:"plan_query_list>inv_unify_plan_top_query,omitempty"`
}

var poolInvUnifyPlanTopQuerys = sync.Pool{
	New: func() any {
		return new(InvUnifyPlanTopQuerys)
	},
}

// GetInvUnifyPlanTopQuerys() 从对象池中获取InvUnifyPlanTopQuerys
func GetInvUnifyPlanTopQuerys() *InvUnifyPlanTopQuerys {
	return poolInvUnifyPlanTopQuerys.Get().(*InvUnifyPlanTopQuerys)
}

// ReleaseInvUnifyPlanTopQuerys 释放InvUnifyPlanTopQuerys
func ReleaseInvUnifyPlanTopQuerys(v *InvUnifyPlanTopQuerys) {
	v.PlanQueryList = v.PlanQueryList[:0]
	poolInvUnifyPlanTopQuerys.Put(v)
}
