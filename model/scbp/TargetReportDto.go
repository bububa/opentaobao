package scbp

// TargetReportDTO 
type TargetReportDTO struct {
    // 返回实体集合
    TargetEffectList   []TargetEffectDTO `json:"target_effect_list,omitempty" xml:"target_effect_list>target_effect_dto,omitempty"`
}
