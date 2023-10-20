package miniappopen

import (
	"sync"
)

// MiniappInstanceAppOnlineDto 结构体
type MiniappInstanceAppOnlineDto struct {
	// 端信息
	Client string `json:"client,omitempty" xml:"client,omitempty"`
	// 错误信息
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMiniappInstanceAppOnlineDto = sync.Pool{
	New: func() any {
		return new(MiniappInstanceAppOnlineDto)
	},
}

// GetMiniappInstanceAppOnlineDto() 从对象池中获取MiniappInstanceAppOnlineDto
func GetMiniappInstanceAppOnlineDto() *MiniappInstanceAppOnlineDto {
	return poolMiniappInstanceAppOnlineDto.Get().(*MiniappInstanceAppOnlineDto)
}

// ReleaseMiniappInstanceAppOnlineDto 释放MiniappInstanceAppOnlineDto
func ReleaseMiniappInstanceAppOnlineDto(v *MiniappInstanceAppOnlineDto) {
	v.Client = ""
	v.FailMsg = ""
	v.Success = ""
	poolMiniappInstanceAppOnlineDto.Put(v)
}
