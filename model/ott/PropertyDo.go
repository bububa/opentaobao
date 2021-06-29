package ott

// PropertyDo 
type PropertyDo struct {
    // 属性键值对
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // 属性校验码
    VerifyCode   string `json:"verify_code,omitempty" xml:"verify_code,omitempty"`
}
