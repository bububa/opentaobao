package axintrade

import (
	"sync"
)

// PackageOtherProductDto 结构体
type PackageOtherProductDto struct {
	// 其他产品名称
	OtherName string `json:"other_name,omitempty" xml:"other_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 其他产品id
	OtherId int64 `json:"other_id,omitempty" xml:"other_id,omitempty"`
}

var poolPackageOtherProductDto = sync.Pool{
	New: func() any {
		return new(PackageOtherProductDto)
	},
}

// GetPackageOtherProductDto() 从对象池中获取PackageOtherProductDto
func GetPackageOtherProductDto() *PackageOtherProductDto {
	return poolPackageOtherProductDto.Get().(*PackageOtherProductDto)
}

// ReleasePackageOtherProductDto 释放PackageOtherProductDto
func ReleasePackageOtherProductDto(v *PackageOtherProductDto) {
	v.OtherName = ""
	v.Remark = ""
	v.OtherId = 0
	poolPackageOtherProductDto.Put(v)
}
