package yunosminiapp

// SystemParam 结构体
type SystemParam struct {
	// 流程id，随机字符串
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 业务码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 小程序id
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 授权token
	AcToken string `json:"ac_token,omitempty" xml:"ac_token,omitempty"`
	// 更新access_token
	ModifyToken bool `json:"modify_token,omitempty" xml:"modify_token,omitempty"`
	// token过期
	TokenExpired bool `json:"token_expired,omitempty" xml:"token_expired,omitempty"`
}
