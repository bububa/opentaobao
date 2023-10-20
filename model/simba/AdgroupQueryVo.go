package simba

import (
	"sync"
)

// AdgroupQueryVo 结构体
type AdgroupQueryVo struct {
	// 计划id集合
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolAdgroupQueryVo = sync.Pool{
	New: func() any {
		return new(AdgroupQueryVo)
	},
}

// GetAdgroupQueryVo() 从对象池中获取AdgroupQueryVo
func GetAdgroupQueryVo() *AdgroupQueryVo {
	return poolAdgroupQueryVo.Get().(*AdgroupQueryVo)
}

// ReleaseAdgroupQueryVo 释放AdgroupQueryVo
func ReleaseAdgroupQueryVo(v *AdgroupQueryVo) {
	v.CampaignIdList = v.CampaignIdList[:0]
	v.Offset = 0
	v.PageSize = 0
	poolAdgroupQueryVo.Put(v)
}
