package wdk

// UserInfoTopDto 结构体
type UserInfoTopDto struct {
	// 人员account
	UserAccount string `json:"user_account,omitempty" xml:"user_account,omitempty"`
	// 人员名字
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}
