package openim

// TribeInfo 
/* model for simplify = false
type TribeInfo struct {

    // 群ID
    
    TribeId   int64 `json:"tribe_id,omitempty"`
    

    // 群头像URL地址
    
    Icon   string `json:"icon,omitempty"`
    

    // 群验证模式
    
    CheckMode   int64 `json:"check_mode,omitempty"`
    

    // 群类型
    
    TribeType   int64 `json:"tribe_type,omitempty"`
    

    // 群名称
    
    Name   string `json:"name,omitempty"`
    

    // 群接收标记
    
    RecvFlag   int64 `json:"recv_flag,omitempty"`
    

    // 群公告
    
    Notice   string `json:"notice,omitempty"`
    

}
*/

// TribeInfo 
type TribeInfo struct {

    // 群ID
    TribeId   int64 `json:"tribe_id,omitempty"`

    // 群头像URL地址
    Icon   string `json:"icon,omitempty"`

    // 群验证模式
    CheckMode   int64 `json:"check_mode,omitempty"`

    // 群类型
    TribeType   int64 `json:"tribe_type,omitempty"`

    // 群名称
    Name   string `json:"name,omitempty"`

    // 群接收标记
    RecvFlag   int64 `json:"recv_flag,omitempty"`

    // 群公告
    Notice   string `json:"notice,omitempty"`

}
