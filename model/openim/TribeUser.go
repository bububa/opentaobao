package openim

// TribeUser 
type TribeUser struct {

    // 用户id
    
    Uid   string `json:"uid,omitempty" xml:"uid,omitempty"`
    

    // 群成员角色
    
    Role   string `json:"role,omitempty" xml:"role,omitempty"`
    

    // 是否为淘宝账号
    
    TaobaoAccount   bool `json:"taobao_account,omitempty" xml:"taobao_account,omitempty"`
    

    // 账户appkey
    
    AppKey   string `json:"app_key,omitempty" xml:"app_key,omitempty"`
    

}
