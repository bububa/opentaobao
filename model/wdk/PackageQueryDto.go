package wdk

import (
	"sync"
)

// PackageQueryDto 结构体
type PackageQueryDto struct {
	// 同城令牌号
	TokenCode string `json:"token_code,omitempty" xml:"token_code,omitempty"`
	// 仓Code
	NodeCode string `json:"node_code,omitempty" xml:"node_code,omitempty"`
	// 作业单元ID
	WorkUnitId string `json:"work_unit_id,omitempty" xml:"work_unit_id,omitempty"`
}

var poolPackageQueryDto = sync.Pool{
	New: func() any {
		return new(PackageQueryDto)
	},
}

// GetPackageQueryDto() 从对象池中获取PackageQueryDto
func GetPackageQueryDto() *PackageQueryDto {
	return poolPackageQueryDto.Get().(*PackageQueryDto)
}

// ReleasePackageQueryDto 释放PackageQueryDto
func ReleasePackageQueryDto(v *PackageQueryDto) {
	v.TokenCode = ""
	v.NodeCode = ""
	v.WorkUnitId = ""
	poolPackageQueryDto.Put(v)
}
