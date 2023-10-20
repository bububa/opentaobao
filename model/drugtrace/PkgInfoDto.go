package drugtrace

import (
	"sync"
)

// PkgInfoDto 结构体
type PkgInfoDto struct {
	// 码信息
	CodeList []string `json:"code_list,omitempty" xml:"code_list>string,omitempty"`
}

var poolPkgInfoDto = sync.Pool{
	New: func() any {
		return new(PkgInfoDto)
	},
}

// GetPkgInfoDto() 从对象池中获取PkgInfoDto
func GetPkgInfoDto() *PkgInfoDto {
	return poolPkgInfoDto.Get().(*PkgInfoDto)
}

// ReleasePkgInfoDto 释放PkgInfoDto
func ReleasePkgInfoDto(v *PkgInfoDto) {
	v.CodeList = v.CodeList[:0]
	poolPkgInfoDto.Put(v)
}
