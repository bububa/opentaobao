package alitripmerchant

// NewUserInfo 
type NewUserInfo struct {

    // 国家
    
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    

    // 性别
    
    Gender   int64 `json:"gender,omitempty" xml:"gender,omitempty"`
    

    // 省份
    
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    

    // 城市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 头像
    
    AvatarUrl   string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
    

    // 昵称
    
    NickName   string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
    

    // 语言
    
    Language   string `json:"language,omitempty" xml:"language,omitempty"`
    

}
