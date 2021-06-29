package scbp

// AccountReportDTO 
type AccountReportDTO struct {
    // 返回数据集合
    AccountEffectList   []AccountEffectDTO `json:"account_effect_list,omitempty" xml:"account_effect_list>account_effect_dto,omitempty"`
}
