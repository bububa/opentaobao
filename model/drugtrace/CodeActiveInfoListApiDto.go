package drugtrace

import (
	"sync"
)

// CodeActiveInfoListApiDto 结构体
type CodeActiveInfoListApiDto struct {
	// 包装比例
	PkgRatio string `json:"pkg_ratio,omitempty" xml:"pkg_ratio,omitempty"`
}

var poolCodeActiveInfoListApiDto = sync.Pool{
	New: func() any {
		return new(CodeActiveInfoListApiDto)
	},
}

// GetCodeActiveInfoListApiDto() 从对象池中获取CodeActiveInfoListApiDto
func GetCodeActiveInfoListApiDto() *CodeActiveInfoListApiDto {
	return poolCodeActiveInfoListApiDto.Get().(*CodeActiveInfoListApiDto)
}

// ReleaseCodeActiveInfoListApiDto 释放CodeActiveInfoListApiDto
func ReleaseCodeActiveInfoListApiDto(v *CodeActiveInfoListApiDto) {
	v.PkgRatio = ""
	poolCodeActiveInfoListApiDto.Put(v)
}
