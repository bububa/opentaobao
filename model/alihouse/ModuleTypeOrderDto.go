package alihouse

import (
	"sync"
)

// ModuleTypeOrderDto 结构体
type ModuleTypeOrderDto struct {
	// 模块ID
	ModuleId string `json:"module_id,omitempty" xml:"module_id,omitempty"`
	// 模块类型
	ModuleType int64 `json:"module_type,omitempty" xml:"module_type,omitempty"`
	// 模块排序
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
}

var poolModuleTypeOrderDto = sync.Pool{
	New: func() any {
		return new(ModuleTypeOrderDto)
	},
}

// GetModuleTypeOrderDto() 从对象池中获取ModuleTypeOrderDto
func GetModuleTypeOrderDto() *ModuleTypeOrderDto {
	return poolModuleTypeOrderDto.Get().(*ModuleTypeOrderDto)
}

// ReleaseModuleTypeOrderDto 释放ModuleTypeOrderDto
func ReleaseModuleTypeOrderDto(v *ModuleTypeOrderDto) {
	v.ModuleId = ""
	v.ModuleType = 0
	v.OrderNo = 0
	poolModuleTypeOrderDto.Put(v)
}
