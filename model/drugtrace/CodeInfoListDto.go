package drugtrace

import (
	"sync"
)

// CodeInfoListDto 结构体
type CodeInfoListDto struct {
	// 制剂规格
	PrepnSpec string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
	// 最小制剂数量
	PrepnAmount string `json:"prepn_amount,omitempty" xml:"prepn_amount,omitempty"`
	// 最小包装数量
	PkgAmount string `json:"pkg_amount,omitempty" xml:"pkg_amount,omitempty"`
	// 监管码级别
	CodeLevel string `json:"code_level,omitempty" xml:"code_level,omitempty"`
	// 监管码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolCodeInfoListDto = sync.Pool{
	New: func() any {
		return new(CodeInfoListDto)
	},
}

// GetCodeInfoListDto() 从对象池中获取CodeInfoListDto
func GetCodeInfoListDto() *CodeInfoListDto {
	return poolCodeInfoListDto.Get().(*CodeInfoListDto)
}

// ReleaseCodeInfoListDto 释放CodeInfoListDto
func ReleaseCodeInfoListDto(v *CodeInfoListDto) {
	v.PrepnSpec = ""
	v.PrepnAmount = ""
	v.PkgAmount = ""
	v.CodeLevel = ""
	v.Code = ""
	poolCodeInfoListDto.Put(v)
}
