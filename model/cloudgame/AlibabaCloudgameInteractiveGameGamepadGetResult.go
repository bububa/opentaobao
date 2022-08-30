package cloudgame

// AlibabaCloudgameInteractiveGameGamepadGetResult 结构体
type AlibabaCloudgameInteractiveGameGamepadGetResult struct {
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据
	Data *GamepadGetResponse `json:"data,omitempty" xml:"data,omitempty"`
}
