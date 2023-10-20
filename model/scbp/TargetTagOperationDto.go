package scbp

import (
	"sync"
)

// TargetTagOperationDto 结构体
type TargetTagOperationDto struct {
	// crowd or region or terminal or scene
	Scope []string `json:"scope,omitempty" xml:"scope>string,omitempty"`
	// 出价类型：0=出价, 1=溢价，2=过滤, 3=召回
	PriceMode int64 `json:"price_mode,omitempty" xml:"price_mode,omitempty"`
}

var poolTargetTagOperationDto = sync.Pool{
	New: func() any {
		return new(TargetTagOperationDto)
	},
}

// GetTargetTagOperationDto() 从对象池中获取TargetTagOperationDto
func GetTargetTagOperationDto() *TargetTagOperationDto {
	return poolTargetTagOperationDto.Get().(*TargetTagOperationDto)
}

// ReleaseTargetTagOperationDto 释放TargetTagOperationDto
func ReleaseTargetTagOperationDto(v *TargetTagOperationDto) {
	v.Scope = v.Scope[:0]
	v.PriceMode = 0
	poolTargetTagOperationDto.Put(v)
}
