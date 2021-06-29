package user

// OpenUserInfoDto 
type OpenUserInfoDto struct {
    // 混淆字符串
    OpenId   string `json:"open_id,omitempty" xml:"open_id,omitempty"`
    // 头像链接
    Avatar   string `json:"avatar,omitempty" xml:"avatar,omitempty"`
    // snsNick
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
}
