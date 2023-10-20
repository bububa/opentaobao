package alilabs

// ScanQrCodeResultVo 结构体
type ScanQrCodeResultVo struct {
	// 天猫精灵设备ID
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 天猫精灵用户ID
	UserOpenId string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
}
