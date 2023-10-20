package tmallnr

import (
	"sync"
)

// TagRespDto 结构体
type TagRespDto struct {
	// 失败描述
	Descs []string `json:"descs,omitempty" xml:"descs>string,omitempty"`
	// 失败商品编码
	FailIds []int64 `json:"fail_ids,omitempty" xml:"fail_ids>int64,omitempty"`
	// 成功商品编码
	SuccessIds []int64 `json:"success_ids,omitempty" xml:"success_ids>int64,omitempty"`
}

var poolTagRespDto = sync.Pool{
	New: func() any {
		return new(TagRespDto)
	},
}

// GetTagRespDto() 从对象池中获取TagRespDto
func GetTagRespDto() *TagRespDto {
	return poolTagRespDto.Get().(*TagRespDto)
}

// ReleaseTagRespDto 释放TagRespDto
func ReleaseTagRespDto(v *TagRespDto) {
	v.Descs = v.Descs[:0]
	v.FailIds = v.FailIds[:0]
	v.SuccessIds = v.SuccessIds[:0]
	poolTagRespDto.Put(v)
}
