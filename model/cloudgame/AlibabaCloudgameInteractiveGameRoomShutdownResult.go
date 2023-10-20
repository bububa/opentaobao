package cloudgame

// AlibabacloudgameinteractivegameroomshutdownResult 结构体
type AlibabacloudgameinteractivegameroomshutdownResult struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 强制停止房间返回结果
	Data *ShutdownRoomResponse `json:"data,omitempty" xml:"data,omitempty"`
}
