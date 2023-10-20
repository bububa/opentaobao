package cloudgame

// AlibabacloudgameinteractivegamestartResult 结构体
type AlibabacloudgameinteractivegamestartResult struct {
	// 返回状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *StartGameResponse `json:"data,omitempty" xml:"data,omitempty"`
}
