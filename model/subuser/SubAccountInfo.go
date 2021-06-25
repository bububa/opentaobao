package subuser

// SubAccountInfo 
type SubAccountInfo struct {

    // 123456
    SubId   int64 `json:"sub_id,omitempty"`

    // zhangsan:no1
    SubNick   string `json:"sub_nick,omitempty"`

    // 654321
    UserId   int64 `json:"user_id,omitempty"`

    // zhangsan
    UserNick   string `json:"user_nick,omitempty"`

    // 1
    SubStatus   int64 `json:"sub_status,omitempty"`

    // true
    SubOwedStatus   bool `json:"sub_owed_status,omitempty"`

    // true
    SubDispatchStatus   bool `json:"sub_dispatch_status,omitempty"`

    // 小红
    SubName   string `json:"sub_name,omitempty"`

}
