package wdk

// UserSyncInfo 
type UserSyncInfo struct {
    // 用户Id
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 手机号
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    // 会员名
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    // 用户昵称
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    // 真实姓名
    RealName   string `json:"real_name,omitempty" xml:"real_name,omitempty"`
    // 性别
    Gender   string `json:"gender,omitempty" xml:"gender,omitempty"`
    // 操作类型 （create/update/delete）
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
