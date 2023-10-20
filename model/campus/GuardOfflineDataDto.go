package campus

import (
	"sync"
)

// GuardOfflineDataDto 结构体
type GuardOfflineDataDto struct {
	// 文件url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// guard
	Guard *GuardDto `json:"guard,omitempty" xml:"guard,omitempty"`
}

var poolGuardOfflineDataDto = sync.Pool{
	New: func() any {
		return new(GuardOfflineDataDto)
	},
}

// GetGuardOfflineDataDto() 从对象池中获取GuardOfflineDataDto
func GetGuardOfflineDataDto() *GuardOfflineDataDto {
	return poolGuardOfflineDataDto.Get().(*GuardOfflineDataDto)
}

// ReleaseGuardOfflineDataDto 释放GuardOfflineDataDto
func ReleaseGuardOfflineDataDto(v *GuardOfflineDataDto) {
	v.Url = ""
	v.Guard = nil
	poolGuardOfflineDataDto.Put(v)
}
