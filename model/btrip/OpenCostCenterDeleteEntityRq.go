package btrip

// OpenCostCenterDeleteEntityRq 
type OpenCostCenterDeleteEntityRq struct {

    // 是否全部删除
    
    DelAll   bool `json:"del_all,omitempty" xml:"del_all,omitempty"`
    

    // 第三方成本中心id
    
    ThirdpartId   string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
    

    // 删除的成员信息列表,del_all为true时可不填
    
    EntityList   []OpenOrgEntityDo `json:"entity_list,omitempty" xml:"entity_list,omitempty"`
    

    // 企业id
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 商旅开放平台传2
    
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    

}
