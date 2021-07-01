package antifraud

// SigAuthenticateContext 结构体
type SigAuthenticateContext struct {
	// 接入密码
	AccessKey string `json:"access_key,omitempty" xml:"access_key,omitempty"`
	// 会话标识,由sip提供的tokenutil工具类生成
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// collinaface JS分配的会话id
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 接入应用采集到的最终用户ip
	RemoteIp string `json:"remote_ip,omitempty" xml:"remote_ip,omitempty"`
	// 签名串
	Sig string `json:"sig,omitempty" xml:"sig,omitempty"`
	// 接入应用标识
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
}
