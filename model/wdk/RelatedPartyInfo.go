package wdk

// RelatedPartyInfo 
type RelatedPartyInfo struct {

    // 所在部门
    
    Department   string `json:"department,omitempty" xml:"department,omitempty"`
    

    // 姓名
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 职位
    
    Post   string `json:"post,omitempty" xml:"post,omitempty"`
    

    // 关联方类型
    
    RelatedPartyType   string `json:"related_party_type,omitempty" xml:"related_party_type,omitempty"`
    

    // 关系
    
    Relation   string `json:"relation,omitempty" xml:"relation,omitempty"`
    

}
