package tblogistics

import (
	"sync"
)

// TPDto 结构体
type TPDto struct {
	// 公司编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 公司名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否商家绑定的默认安装公司
	IsDefault bool `json:"is_default,omitempty" xml:"is_default,omitempty"`
}

var poolTPDto = sync.Pool{
	New: func() any {
		return new(TPDto)
	},
}

// GetTPDto() 从对象池中获取TPDto
func GetTPDto() *TPDto {
	return poolTPDto.Get().(*TPDto)
}

// ReleaseTPDto 释放TPDto
func ReleaseTPDto(v *TPDto) {
	v.Code = ""
	v.Name = ""
	v.IsDefault = false
	poolTPDto.Put(v)
}
