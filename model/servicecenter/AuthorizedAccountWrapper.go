package servicecenter

// AuthorizedAccountWrapper 
type AuthorizedAccountWrapper struct {

    // 商家子账号记录数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    

    // 商家子账号列表
    
    SubUsers   []SubUser `json:"sub_users,omitempty" xml:"sub_users,omitempty"`
    

}
