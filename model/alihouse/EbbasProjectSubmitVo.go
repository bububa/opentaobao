package alihouse

import (
	"sync"
)

// EbbasProjectSubmitVo 结构体
type EbbasProjectSubmitVo struct {
	// 楼盘code
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 楼盘e码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}

var poolEbbasProjectSubmitVo = sync.Pool{
	New: func() any {
		return new(EbbasProjectSubmitVo)
	},
}

// GetEbbasProjectSubmitVo() 从对象池中获取EbbasProjectSubmitVo
func GetEbbasProjectSubmitVo() *EbbasProjectSubmitVo {
	return poolEbbasProjectSubmitVo.Get().(*EbbasProjectSubmitVo)
}

// ReleaseEbbasProjectSubmitVo 释放EbbasProjectSubmitVo
func ReleaseEbbasProjectSubmitVo(v *EbbasProjectSubmitVo) {
	v.ProjectCode = ""
	v.ECode = ""
	v.OuterId = ""
	poolEbbasProjectSubmitVo.Put(v)
}
