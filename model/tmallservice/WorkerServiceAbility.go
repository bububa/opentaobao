package tmallservice

import (
	"sync"
)

// WorkerServiceAbility 结构体
type WorkerServiceAbility struct {
	// 服务区域
	AreaCodeList []int64 `json:"area_code_list,omitempty" xml:"area_code_list>int64,omitempty"`
	// 技能
	ServiceSkillList []string `json:"service_skill_list,omitempty" xml:"service_skill_list>string,omitempty"`
	// 身份证
	WorkerId int64 `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
}

var poolWorkerServiceAbility = sync.Pool{
	New: func() any {
		return new(WorkerServiceAbility)
	},
}

// GetWorkerServiceAbility() 从对象池中获取WorkerServiceAbility
func GetWorkerServiceAbility() *WorkerServiceAbility {
	return poolWorkerServiceAbility.Get().(*WorkerServiceAbility)
}

// ReleaseWorkerServiceAbility 释放WorkerServiceAbility
func ReleaseWorkerServiceAbility(v *WorkerServiceAbility) {
	v.AreaCodeList = v.AreaCodeList[:0]
	v.ServiceSkillList = v.ServiceSkillList[:0]
	v.WorkerId = 0
	poolWorkerServiceAbility.Put(v)
}
