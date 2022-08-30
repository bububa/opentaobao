package alitripmerchant

// UserCurrentStatus 结构体
type UserCurrentStatus struct {
	// 用户当前状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
}
