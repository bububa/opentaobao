package btrip

// ClientInfoDo 
type ClientInfoDo struct {
    // 乘机人id
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 乘机人姓名
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}
