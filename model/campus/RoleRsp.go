package campus

// RoleRsp 
type RoleRsp struct {

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误级别
    
    ErrorLevel   string `json:"error_level,omitempty" xml:"error_level,omitempty"`
    

    // 角色id
    
    RoleId   string `json:"role_id,omitempty" xml:"role_id,omitempty"`
    

}
