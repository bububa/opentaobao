package servicecenter

// AuthorizedAccountWrapper 
/* model for simplify = false
type AuthorizedAccountWrapper struct {

    // 商家子账号记录数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 商家子账号列表
    
    SubUsers  struct {
        SubUser  []SubUser `json:"sub_user,omitempty"`
    } `json:"sub_users,omitempty"`
    

}
*/

// AuthorizedAccountWrapper 
type AuthorizedAccountWrapper struct {

    // 商家子账号记录数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 商家子账号列表
    SubUsers   []SubUser `json:"sub_users,omitempty"`

}
