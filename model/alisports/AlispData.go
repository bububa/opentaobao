package alisports

// AlispData 
type AlispData struct {
    // sso_token
    SsoToken   string `json:"sso_token,omitempty" xml:"sso_token,omitempty"`
    // aliuid
    Aliuid   string `json:"aliuid,omitempty" xml:"aliuid,omitempty"`
    // type
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // access_token
    AccessToken   string `json:"access_token,omitempty" xml:"access_token,omitempty"`
    // 头像
    AvatarUrl   string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
    // 昵称
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
    // 手机号，需要找阿里体育申请并且配置白名单返回
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}
