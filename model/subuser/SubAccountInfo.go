package subuser

// SubAccountInfo 结构体
type SubAccountInfo struct {
	// zhangsan:no1
	SubNick string `json:"sub_nick,omitempty" xml:"sub_nick,omitempty"`
	// zhangsan
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 小红
	SubName string `json:"sub_name,omitempty" xml:"sub_name,omitempty"`
	// 123456
	SubId int64 `json:"sub_id,omitempty" xml:"sub_id,omitempty"`
	// 654321
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 1
	SubStatus int64 `json:"sub_status,omitempty" xml:"sub_status,omitempty"`
	// true
	SubOwedStatus bool `json:"sub_owed_status,omitempty" xml:"sub_owed_status,omitempty"`
	// true
	SubDispatchStatus bool `json:"sub_dispatch_status,omitempty" xml:"sub_dispatch_status,omitempty"`
}
