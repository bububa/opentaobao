package media

import (
	"sync"
)

// TopUserStorageDto 结构体
type TopUserStorageDto struct {
	// 已使用的容量，单位字节
	UsedQuota int64 `json:"used_quota,omitempty" xml:"used_quota,omitempty"`
	// 总容量，单位字节
	TotalQuota int64 `json:"total_quota,omitempty" xml:"total_quota,omitempty"`
}

var poolTopUserStorageDto = sync.Pool{
	New: func() any {
		return new(TopUserStorageDto)
	},
}

// GetTopUserStorageDto() 从对象池中获取TopUserStorageDto
func GetTopUserStorageDto() *TopUserStorageDto {
	return poolTopUserStorageDto.Get().(*TopUserStorageDto)
}

// ReleaseTopUserStorageDto 释放TopUserStorageDto
func ReleaseTopUserStorageDto(v *TopUserStorageDto) {
	v.UsedQuota = 0
	v.TotalQuota = 0
	poolTopUserStorageDto.Put(v)
}
