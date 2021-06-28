package openim

// User 
/* model for simplify = false
type User struct {

    // 用户所属appkey
    
    Appkey   string `json:"appkey,omitempty"`
    

    // 是否是淘宝账号
    
    TaobaoAccount   bool `json:"taobao_account,omitempty"`
    

    // 用户id
    
    Uid   string `json:"uid,omitempty"`
    

    // 账户appkey
    
    AppKey   string `json:"app_key,omitempty"`
    

}
*/

// User 
type User struct {

    // 用户所属appkey
    Appkey   string `json:"appkey,omitempty"`

    // 是否是淘宝账号
    TaobaoAccount   bool `json:"taobao_account,omitempty"`

    // 用户id
    Uid   string `json:"uid,omitempty"`

    // 账户appkey
    AppKey   string `json:"app_key,omitempty"`

}
