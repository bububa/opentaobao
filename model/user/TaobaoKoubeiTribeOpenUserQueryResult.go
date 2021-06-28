package user

// TaobaoKoubeiTribeOpenUserQueryResult 
/* model for simplify = false
type TaobaoKoubeiTribeOpenUserQueryResult struct {

    // traceId
    
    TraceId   string `json:"trace_id,omitempty"`
    

    // 用户信息DTO
    
    Data  *struct {
        UserInfoDto  *UserInfoDto `json:"user_info_dto,omitempty"`
    } `json:"data,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误信息
    
    Error   string `json:"error,omitempty"`
    

}
*/

// TaobaoKoubeiTribeOpenUserQueryResult 
type TaobaoKoubeiTribeOpenUserQueryResult struct {

    // traceId
    TraceId   string `json:"trace_id,omitempty"`

    // 用户信息DTO
    Data   *UserInfoDto `json:"data,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误信息
    Error   string `json:"error,omitempty"`

}
