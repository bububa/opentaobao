package scbp

// TargetReportDto 
/* model for simplify = false
type TargetReportDto struct {

    // 返回实体集合
    
    TargetEffectList  struct {
        TargetEffectDto  []TargetEffectDto `json:"target_effect_dto,omitempty"`
    } `json:"target_effect_list,omitempty"`
    

}
*/

// TargetReportDto 
type TargetReportDto struct {

    // 返回实体集合
    TargetEffectList   []TargetEffectDto `json:"target_effect_list,omitempty"`

}
