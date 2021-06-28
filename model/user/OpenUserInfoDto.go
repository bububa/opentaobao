package user

// OpenUserInfoDto 
/* model for simplify = false
type OpenUserInfoDto struct {

    // 混淆字符串
    
    OpenId   string `json:"open_id,omitempty"`
    

    // 头像链接
    
    Avatar   string `json:"avatar,omitempty"`
    

    // snsNick
    
    Nick   string `json:"nick,omitempty"`
    

}
*/

// OpenUserInfoDto 
type OpenUserInfoDto struct {

    // 混淆字符串
    OpenId   string `json:"open_id,omitempty"`

    // 头像链接
    Avatar   string `json:"avatar,omitempty"`

    // snsNick
    Nick   string `json:"nick,omitempty"`

}
