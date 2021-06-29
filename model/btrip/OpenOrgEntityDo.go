package btrip

// OpenOrgEntityDo 
type OpenOrgEntityDo struct {

    // 用户/部门/角色名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 用户/部门/角色id
    
    EntityId   string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
    

    // 人员类型：1用户，2部门，3角色
    
    EntityType   string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
    

    // 企业id
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 角色/部门下面员工人数
    
    UserNum   int64 `json:"user_num,omitempty" xml:"user_num,omitempty"`
    

}
