package user

// UserInfoBaseDto 
type UserInfoBaseDto struct {

    // 同一个开放平台帐号下，用户的 UnionID 是唯一的
    
    UnionId   string `json:"union_id,omitempty" xml:"union_id,omitempty"`
    

    // 小程序用户唯一识别
    
    OpenUid   string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
    

    // 小程序ID
    
    AppId   string `json:"app_id,omitempty" xml:"app_id,omitempty"`
    

    // 小程序开发主体ID
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 登录类型
    
    LoginType   string `json:"login_type,omitempty" xml:"login_type,omitempty"`
    

    // 不同业务/登录方式的扩展字段
    
    ExtraInfo   string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
    

    // 登录用户昵称
    
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    

}
