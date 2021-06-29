package btrip

// OpenCostCenterQueryRq 
type OpenCostCenterQueryRq struct {

    // 第三方成本中心id，不填写的时候user_id必填
    
    ThirdpartId   string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
    

    // 用户id，不填的时候thirdpart_id必填
    
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 是否需要展示成员信息，当成本中心为部分人员适用的时候有返回
    
    NeedOrgEntity   bool `json:"need_org_entity,omitempty" xml:"need_org_entity,omitempty"`
    

    // 企业id
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 成本中心名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 商旅开放平台传2
    
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    

}
