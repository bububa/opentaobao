package axintrade

import (
	"sync"
)

// PackageApiResDto 结构体
type PackageApiResDto struct {
	// 产品id列表
	ProductIds []int64 `json:"product_ids,omitempty" xml:"product_ids>int64,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页数
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPackageApiResDto = sync.Pool{
	New: func() any {
		return new(PackageApiResDto)
	},
}

// GetPackageApiResDto() 从对象池中获取PackageApiResDto
func GetPackageApiResDto() *PackageApiResDto {
	return poolPackageApiResDto.Get().(*PackageApiResDto)
}

// ReleasePackageApiResDto 释放PackageApiResDto
func ReleasePackageApiResDto(v *PackageApiResDto) {
	v.ProductIds = v.ProductIds[:0]
	v.TotalCount = 0
	v.PageNo = 0
	v.PageSize = 0
	poolPackageApiResDto.Put(v)
}
