package btrip

// OpenCostCenterSetEntityRq 
type OpenCostCenterSetEntityRq struct {
    // 第三方成本中心id
    ThirdpartId   string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
    // 人员信息列表
    EntityList   []OpenOrgEntityDo `json:"entity_list,omitempty" xml:"entity_list>open_org_entity_do,omitempty"`
    // 企业id
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 商旅开放平台传2
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
}
