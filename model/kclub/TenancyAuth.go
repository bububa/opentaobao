package kclub

// TenancyAuth 
type TenancyAuth struct {
    // 鉴权秘钥
    SecretKey   string `json:"secret_key,omitempty" xml:"secret_key,omitempty"`
    // 鉴权应用名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
