package drugtrace

import (
	"sync"
)

// CodeScanDto 结构体
type CodeScanDto struct {
	// 扫码时间
	QueryDate string `json:"query_date,omitempty" xml:"query_date,omitempty"`
	// 扫码来源：支付宝，淘宝，天猫，未知
	QueryAppNameFormat string `json:"query_app_name_format,omitempty" xml:"query_app_name_format,omitempty"`
}

var poolCodeScanDto = sync.Pool{
	New: func() any {
		return new(CodeScanDto)
	},
}

// GetCodeScanDto() 从对象池中获取CodeScanDto
func GetCodeScanDto() *CodeScanDto {
	return poolCodeScanDto.Get().(*CodeScanDto)
}

// ReleaseCodeScanDto 释放CodeScanDto
func ReleaseCodeScanDto(v *CodeScanDto) {
	v.QueryDate = ""
	v.QueryAppNameFormat = ""
	poolCodeScanDto.Put(v)
}
