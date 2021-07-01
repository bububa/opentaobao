package scbp

// TargetReportDto 结构体
type TargetReportDto struct {
	// 返回实体集合
	TargetEffectList []TargetEffectDto `json:"target_effect_list,omitempty" xml:"target_effect_list>target_effect_dto,omitempty"`
}
