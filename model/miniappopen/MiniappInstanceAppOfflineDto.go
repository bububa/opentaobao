package miniappopen

import (
	"sync"
)

// MiniappInstanceAppOfflineDto 结构体
type MiniappInstanceAppOfflineDto struct {
	// 端信息
	Client string `json:"client,omitempty" xml:"client,omitempty"`
	// 错误信息
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMiniappInstanceAppOfflineDto = sync.Pool{
	New: func() any {
		return new(MiniappInstanceAppOfflineDto)
	},
}

// GetMiniappInstanceAppOfflineDto() 从对象池中获取MiniappInstanceAppOfflineDto
func GetMiniappInstanceAppOfflineDto() *MiniappInstanceAppOfflineDto {
	return poolMiniappInstanceAppOfflineDto.Get().(*MiniappInstanceAppOfflineDto)
}

// ReleaseMiniappInstanceAppOfflineDto 释放MiniappInstanceAppOfflineDto
func ReleaseMiniappInstanceAppOfflineDto(v *MiniappInstanceAppOfflineDto) {
	v.Client = ""
	v.FailMsg = ""
	v.Success = false
	poolMiniappInstanceAppOfflineDto.Put(v)
}
