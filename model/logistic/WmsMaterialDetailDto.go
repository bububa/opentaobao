package logistic

import (
	"sync"
)

// WmsMaterialDetailDto 结构体
type WmsMaterialDetailDto struct {
	// 包材名称
	MaterialName string `json:"material_name,omitempty" xml:"material_name,omitempty"`
	// 包材编码
	MaterialCode string `json:"material_code,omitempty" xml:"material_code,omitempty"`
	// 包材使用数量
	MaterialNum int64 `json:"material_num,omitempty" xml:"material_num,omitempty"`
}

var poolWmsMaterialDetailDto = sync.Pool{
	New: func() any {
		return new(WmsMaterialDetailDto)
	},
}

// GetWmsMaterialDetailDto() 从对象池中获取WmsMaterialDetailDto
func GetWmsMaterialDetailDto() *WmsMaterialDetailDto {
	return poolWmsMaterialDetailDto.Get().(*WmsMaterialDetailDto)
}

// ReleaseWmsMaterialDetailDto 释放WmsMaterialDetailDto
func ReleaseWmsMaterialDetailDto(v *WmsMaterialDetailDto) {
	v.MaterialName = ""
	v.MaterialCode = ""
	v.MaterialNum = 0
	poolWmsMaterialDetailDto.Put(v)
}
