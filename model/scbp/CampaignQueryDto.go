package scbp

// CampaignQueryDto 
type CampaignQueryDto struct {

    // 计划标题，配合exactMatch使用
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 子类型
    
    SubType   string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
    

    // 推广状态
    
    OnlineStatus   int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
    

    // 标题是否精确匹配
    
    ExactMatch   bool `json:"exact_match,omitempty" xml:"exact_match,omitempty"`
    

    // 计划类型列表
    
    TypeList   []int64 `json:"type_list,omitempty" xml:"type_list>int64,omitempty"`
    

    // 计划子类型列表
    
    SubTypeList   []string `json:"sub_type_list,omitempty" xml:"sub_type_list>string,omitempty"`
    

    // 类目id
    
    CateId   int64 `json:"cate_id,omitempty" xml:"cate_id,omitempty"`
    

    // 方案包配置id列表
    
    PkgCfgIdList   []int64 `json:"pkg_cfg_id_list,omitempty" xml:"pkg_cfg_id_list>int64,omitempty"`
    

    // 当前页数
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

    // 每页数量
    
    Size   int64 `json:"size,omitempty" xml:"size,omitempty"`
    

    // 计划id集合
    
    CampaignIdList   []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
    

}
