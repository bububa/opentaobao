package antifraud

// ParamAccountQuery 结构体
type ParamAccountQuery struct {
	// 反欺诈服务AppKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 时间戳
	Timestamp string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 身份校验信息
	AppToken string `json:"app_token,omitempty" xml:"app_token,omitempty"`
	// 场景ID
	SceneId string `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	// 手机号
	PhoneNumber string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// IP
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
	// 透传参数，是一个json形式的字符串
	Trans string `json:"trans,omitempty" xml:"trans,omitempty"`
}
