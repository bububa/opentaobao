package promotion

// BenefitSelectorVo 结构体
type BenefitSelectorVo struct {
	// 详情list
	PackDetailList []BenefitTemplateVo `json:"pack_detail_list,omitempty" xml:"pack_detail_list>benefit_template_vo,omitempty"`
	// 权益类型
	BenefitType string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
	// 权益名称
	BenefitName string `json:"benefit_name,omitempty" xml:"benefit_name,omitempty"`
	// 权益类型id
	BenefitTypeLong int64 `json:"benefit_type_long,omitempty" xml:"benefit_type_long,omitempty"`
}
