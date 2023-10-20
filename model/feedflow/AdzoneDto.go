package feedflow

import (
	"sync"
)

// AdzoneDto 结构体
type AdzoneDto struct {
	// 广告位名称
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
	// 广告位id
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
}

var poolAdzoneDto = sync.Pool{
	New: func() any {
		return new(AdzoneDto)
	},
}

// GetAdzoneDto() 从对象池中获取AdzoneDto
func GetAdzoneDto() *AdzoneDto {
	return poolAdzoneDto.Get().(*AdzoneDto)
}

// ReleaseAdzoneDto 释放AdzoneDto
func ReleaseAdzoneDto(v *AdzoneDto) {
	v.AdzoneName = ""
	v.AdzoneId = 0
	poolAdzoneDto.Put(v)
}
