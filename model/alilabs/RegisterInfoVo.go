package alilabs

// RegisterInfoVo 结构体
type RegisterInfoVo struct {
	// 用户开放id
	UserOpenId string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
	// 设备uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}
