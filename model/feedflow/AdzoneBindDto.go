package feedflow

import (
	"sync"
)

// AdzoneBindDto 结构体
type AdzoneBindDto struct {
	// 广告位名称
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
	// 资源位id
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
	// 溢价
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}

var poolAdzoneBindDto = sync.Pool{
	New: func() any {
		return new(AdzoneBindDto)
	},
}

// GetAdzoneBindDto() 从对象池中获取AdzoneBindDto
func GetAdzoneBindDto() *AdzoneBindDto {
	return poolAdzoneBindDto.Get().(*AdzoneBindDto)
}

// ReleaseAdzoneBindDto 释放AdzoneBindDto
func ReleaseAdzoneBindDto(v *AdzoneBindDto) {
	v.AdzoneName = ""
	v.AdzoneId = 0
	v.Discount = 0
	poolAdzoneBindDto.Put(v)
}
