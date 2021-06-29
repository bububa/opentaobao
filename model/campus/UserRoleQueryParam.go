package campus

// UserRoleQueryParam 
type UserRoleQueryParam struct {

    // 多应用
    
    AppKeys   []string `json:"app_keys,omitempty" xml:"app_keys>string,omitempty"`
    

    // 用户账号
    
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 是否支持有效期过滤
    
    ReturnNotEffective   bool `json:"return_not_effective,omitempty" xml:"return_not_effective,omitempty"`
    

}
