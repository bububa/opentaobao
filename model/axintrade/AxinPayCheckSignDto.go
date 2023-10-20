package axintrade

import (
	"sync"
)

// AxinPayCheckSignDto 结构体
type AxinPayCheckSignDto struct {
	// 业务场景
	BizScene string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
	// 验签数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAxinPayCheckSignDto = sync.Pool{
	New: func() any {
		return new(AxinPayCheckSignDto)
	},
}

// GetAxinPayCheckSignDto() 从对象池中获取AxinPayCheckSignDto
func GetAxinPayCheckSignDto() *AxinPayCheckSignDto {
	return poolAxinPayCheckSignDto.Get().(*AxinPayCheckSignDto)
}

// ReleaseAxinPayCheckSignDto 释放AxinPayCheckSignDto
func ReleaseAxinPayCheckSignDto(v *AxinPayCheckSignDto) {
	v.BizScene = ""
	v.Data = ""
	poolAxinPayCheckSignDto.Put(v)
}
