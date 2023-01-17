package cloudgame

// HavanaUserIdQueryRequest 结构体
type HavanaUserIdQueryRequest struct {
	// API名称
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 前端APPKEY
	FrontAppKey string `json:"front_app_key,omitempty" xml:"front_app_key,omitempty"`
	// 登录态token
	AccountToken string `json:"account_token,omitempty" xml:"account_token,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
