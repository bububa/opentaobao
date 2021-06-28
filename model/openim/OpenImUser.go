package openim

// OpenImUser 
/* model for simplify = false
type OpenImUser struct {

    // 用户id
    
    Uid   string `json:"uid,omitempty"`
    

    // 是否为淘宝账号
    
    TaobaoAccount   bool `json:"taobao_account,omitempty"`
    

    // 账户appkey
    
    AppKey   string `json:"app_key,omitempty"`
    

}
*/

// OpenImUser 
type OpenImUser struct {

    // 用户id
    Uid   string `json:"uid,omitempty"`

    // 是否为淘宝账号
    TaobaoAccount   bool `json:"taobao_account,omitempty"`

    // 账户appkey
    AppKey   string `json:"app_key,omitempty"`

}
