package cloudgame

// AlibabaCloudgameInteractiveGameRoomCreateResult 结构体
type AlibabaCloudgameInteractiveGameRoomCreateResult struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *CreateRoomResponse `json:"data,omitempty" xml:"data,omitempty"`
}
