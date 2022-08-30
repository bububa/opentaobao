package tmallgenie

// AuthResultVo 结构体
type AuthResultVo struct {
	// 设备uuid
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 精灵用户openid
	UserOpenId string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
}
