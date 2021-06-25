package promotion

// BenefitSelectorVo 
type BenefitSelectorVo struct {

    // 权益类型
    BenefitType   string `json:"benefit_type,omitempty"`

    // 详情list
    PackDetailList   []BenefitTemplateVo `json:"pack_detail_list,omitempty"`

    // 权益名称
    BenefitName   string `json:"benefit_name,omitempty"`

    // 权益类型id
    BenefitTypeLong   int64 `json:"benefit_type_long,omitempty"`

}
