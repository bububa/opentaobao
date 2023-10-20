package yunosminiapp

// YunosminiappdatatunnelcallBaseRequest 结构体
type YunosminiappdatatunnelcallBaseRequest struct {
	// 请求基础参数
	SystemParam *SystemParam `json:"system_param,omitempty" xml:"system_param,omitempty"`
	// 请求参数
	BizParam *YunosminiappdatatunnelcallBizParam `json:"biz_param,omitempty" xml:"biz_param,omitempty"`
}
