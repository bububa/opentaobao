package alisports

// CheckEfspUserVo 
type CheckEfspUserVo struct {
    // 健身房IDS
    GymList   string `json:"gym_list,omitempty" xml:"gym_list,omitempty"`
    // 剩余补助
    SurplusAmount   int64 `json:"surplus_amount,omitempty" xml:"surplus_amount,omitempty"`
    // 是否企业健身用户
    IsEfsp   bool `json:"is_efsp,omitempty" xml:"is_efsp,omitempty"`
    // 公司名称
    EnterpriseName   string `json:"enterprise_name,omitempty" xml:"enterprise_name,omitempty"`
    // 公司code
    EnterpriseId   string `json:"enterprise_id,omitempty" xml:"enterprise_id,omitempty"`
    // 规则信息
    Rule   *RuleVo `json:"rule,omitempty" xml:"rule,omitempty"`
    // 健身房黑名单
    GymIdsNotSelect   string `json:"gym_ids_not_select,omitempty" xml:"gym_ids_not_select,omitempty"`
    // 健身房地区白名单
    CityCodes   string `json:"city_codes,omitempty" xml:"city_codes,omitempty"`
}
