package tvupadmin

import (
	"sync"
)

// AppDto 结构体
type AppDto struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// packageName
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 应用ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolAppDto = sync.Pool{
	New: func() any {
		return new(AppDto)
	},
}

// GetAppDto() 从对象池中获取AppDto
func GetAppDto() *AppDto {
	return poolAppDto.Get().(*AppDto)
}

// ReleaseAppDto 释放AppDto
func ReleaseAppDto(v *AppDto) {
	v.Name = ""
	v.PackageName = ""
	v.Id = 0
	poolAppDto.Put(v)
}
