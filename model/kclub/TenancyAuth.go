package kclub

// TenancyAuth 
/* model for simplify = false
type TenancyAuth struct {

    // 鉴权秘钥
    
    SecretKey   string `json:"secret_key,omitempty"`
    

    // 鉴权应用名称
    
    Name   string `json:"name,omitempty"`
    

}
*/

// TenancyAuth 
type TenancyAuth struct {

    // 鉴权秘钥
    SecretKey   string `json:"secret_key,omitempty"`

    // 鉴权应用名称
    Name   string `json:"name,omitempty"`

}
