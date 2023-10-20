package simba

import (
	"sync"
)

// WordPackageQueryVo 结构体
type WordPackageQueryVo struct {
	// 计划id集合,计划已经存在场景必填
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 单元id集合,单元已经存在场景必填
	AdgroupIdList []int64 `json:"adgroup_id_list,omitempty" xml:"adgroup_id_list>int64,omitempty"`
}

var poolWordPackageQueryVo = sync.Pool{
	New: func() any {
		return new(WordPackageQueryVo)
	},
}

// GetWordPackageQueryVo() 从对象池中获取WordPackageQueryVo
func GetWordPackageQueryVo() *WordPackageQueryVo {
	return poolWordPackageQueryVo.Get().(*WordPackageQueryVo)
}

// ReleaseWordPackageQueryVo 释放WordPackageQueryVo
func ReleaseWordPackageQueryVo(v *WordPackageQueryVo) {
	v.CampaignIdList = v.CampaignIdList[:0]
	v.AdgroupIdList = v.AdgroupIdList[:0]
	poolWordPackageQueryVo.Put(v)
}
