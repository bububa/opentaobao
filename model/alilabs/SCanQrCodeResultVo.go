package alilabs

import (
	"sync"
)

// SCanQrCodeResultVo 结构体
type SCanQrCodeResultVo struct {
	// 天猫精灵设备ID
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 天猫精灵用户ID
	UserOpenId string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
}

var poolSCanQrCodeResultVo = sync.Pool{
	New: func() any {
		return new(SCanQrCodeResultVo)
	},
}

// GetSCanQrCodeResultVo() 从对象池中获取SCanQrCodeResultVo
func GetSCanQrCodeResultVo() *SCanQrCodeResultVo {
	return poolSCanQrCodeResultVo.Get().(*SCanQrCodeResultVo)
}

// ReleaseSCanQrCodeResultVo 释放SCanQrCodeResultVo
func ReleaseSCanQrCodeResultVo(v *SCanQrCodeResultVo) {
	v.Uuid = ""
	v.UserOpenId = ""
	poolSCanQrCodeResultVo.Put(v)
}
