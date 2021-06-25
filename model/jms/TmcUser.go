package jms

// TmcUser 
type TmcUser struct {

    // 用户昵称
    UserNick   string `json:"user_nick,omitempty"`

    // 用户ID
    UserId   int64 `json:"user_id,omitempty"`

    // 用户授权是否有效，true表示授权有效，false表示授权过期
    IsValid   bool `json:"is_valid,omitempty"`

    // 用户首次开通时间
    Created   string `json:"created,omitempty"`

    // 用户最后开通时间
    Modified   string `json:"modified,omitempty"`

}
