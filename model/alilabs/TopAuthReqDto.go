package alilabs

// TopAuthReqDto 结构体
type TopAuthReqDto struct {
	// 二维码授权 只支持qrcode
	ResponseType string `json:"response_type,omitempty" xml:"response_type,omitempty"`
	// 天猫精灵分配的clientId
	ClientId string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// OAUTH2 scope 只支持basic
	Scope string `json:"scope,omitempty" xml:"scope,omitempty"`
	// OAUTH2 state 随意填写
	State string `json:"state,omitempty" xml:"state,omitempty"`
}
