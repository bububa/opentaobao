package alihouse

import (
	"sync"
)

// EbbasCommunitySubmitVo 结构体
type EbbasCommunitySubmitVo struct {
	// 楼盘code
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 楼盘e码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}

var poolEbbasCommunitySubmitVo = sync.Pool{
	New: func() any {
		return new(EbbasCommunitySubmitVo)
	},
}

// GetEbbasCommunitySubmitVo() 从对象池中获取EbbasCommunitySubmitVo
func GetEbbasCommunitySubmitVo() *EbbasCommunitySubmitVo {
	return poolEbbasCommunitySubmitVo.Get().(*EbbasCommunitySubmitVo)
}

// ReleaseEbbasCommunitySubmitVo 释放EbbasCommunitySubmitVo
func ReleaseEbbasCommunitySubmitVo(v *EbbasCommunitySubmitVo) {
	v.ProjectCode = ""
	v.ECode = ""
	v.OuterId = ""
	poolEbbasCommunitySubmitVo.Put(v)
}
