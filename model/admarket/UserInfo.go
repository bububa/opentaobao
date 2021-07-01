package admarket

// UserInfo 结构体
type UserInfo struct {
	// 用户id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 补充信息
	Info string `json:"info,omitempty" xml:"info,omitempty"`
}
