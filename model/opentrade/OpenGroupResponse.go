package opentrade

// OpenGroupResponse 
type OpenGroupResponse struct {

    // 团id
    
    GroupId   int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
    

    // 成团过期时间，单位为秒
    
    Expiration   int64 `json:"expiration,omitempty" xml:"expiration,omitempty"`
    

    // 成团目标数
    
    Goal   int64 `json:"goal,omitempty" xml:"goal,omitempty"`
    

    // 组团类型，0 拼团 1 成团
    
    GroupType   int64 `json:"group_type,omitempty" xml:"group_type,omitempty"`
    

}
