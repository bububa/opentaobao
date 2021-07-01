package scbp

// AccountReportDto 结构体
type AccountReportDto struct {
	// 返回数据集合
	AccountEffectList []AccountEffectDto `json:"account_effect_list,omitempty" xml:"account_effect_list>account_effect_dto,omitempty"`
}
