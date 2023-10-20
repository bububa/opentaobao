package simba

import (
	"sync"
)

// WordQueryVo 结构体
type WordQueryVo struct {
	// 计划id集合,计划已经存在场景必填
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 单元id集合,单元已经存在场景必填
	AdgroupIdList []int64 `json:"adgroup_id_list,omitempty" xml:"adgroup_id_list>int64,omitempty"`
}

var poolWordQueryVo = sync.Pool{
	New: func() any {
		return new(WordQueryVo)
	},
}

// GetWordQueryVo() 从对象池中获取WordQueryVo
func GetWordQueryVo() *WordQueryVo {
	return poolWordQueryVo.Get().(*WordQueryVo)
}

// ReleaseWordQueryVo 释放WordQueryVo
func ReleaseWordQueryVo(v *WordQueryVo) {
	v.CampaignIdList = v.CampaignIdList[:0]
	v.AdgroupIdList = v.AdgroupIdList[:0]
	poolWordQueryVo.Put(v)
}
