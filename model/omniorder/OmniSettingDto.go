package omniorder

import (
	"sync"
)

// OmniSettingDto 结构体
type OmniSettingDto struct {
	// 分单系统，填 astrolabe 代表阿里分单，填 RDS的 appkey 代表自行分单
	AllocatedSystem string `json:"allocated_system,omitempty" xml:"allocated_system,omitempty"`
	// 接单系统，填 0 代表店掌柜，填 1 代表 POS
	AcceptedSystem int64 `json:"accepted_system,omitempty" xml:"accepted_system,omitempty"`
}

var poolOmniSettingDto = sync.Pool{
	New: func() any {
		return new(OmniSettingDto)
	},
}

// GetOmniSettingDto() 从对象池中获取OmniSettingDto
func GetOmniSettingDto() *OmniSettingDto {
	return poolOmniSettingDto.Get().(*OmniSettingDto)
}

// ReleaseOmniSettingDto 释放OmniSettingDto
func ReleaseOmniSettingDto(v *OmniSettingDto) {
	v.AllocatedSystem = ""
	v.AcceptedSystem = 0
	poolOmniSettingDto.Put(v)
}
