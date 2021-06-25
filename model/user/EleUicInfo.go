package user

// EleUicInfo 
type EleUicInfo struct {

    // 用户头像
    Avatar   string `json:"avatar,omitempty"`

    // 用户昵称
    Nick   string `json:"nick,omitempty"`

    // 手机号
    Phone   string `json:"phone,omitempty"`

    // openId
    OpenId   string `json:"open_id,omitempty"`

}
