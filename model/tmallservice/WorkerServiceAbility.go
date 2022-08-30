package tmallservice

// WorkerServiceAbility 结构体
type WorkerServiceAbility struct {
	// 服务区域
	AreaCodeList []int64 `json:"area_code_list,omitempty" xml:"area_code_list>int64,omitempty"`
	// 技能
	ServiceSkillList []string `json:"service_skill_list,omitempty" xml:"service_skill_list>string,omitempty"`
	// 身份证
	WorkerId int64 `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
}
