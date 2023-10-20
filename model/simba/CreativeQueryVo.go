package simba

import (
	"sync"
)

// CreativeQueryVo 结构体
type CreativeQueryVo struct {
	// 单元id集合
	AdgroupIdList []int64 `json:"adgroup_id_list,omitempty" xml:"adgroup_id_list>int64,omitempty"`
	// 计划id集合
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 报表查询条件
	RptQuery *RptQueryVo `json:"rpt_query,omitempty" xml:"rpt_query,omitempty"`
	// 创意id
	CreativeId int64 `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
}

var poolCreativeQueryVo = sync.Pool{
	New: func() any {
		return new(CreativeQueryVo)
	},
}

// GetCreativeQueryVo() 从对象池中获取CreativeQueryVo
func GetCreativeQueryVo() *CreativeQueryVo {
	return poolCreativeQueryVo.Get().(*CreativeQueryVo)
}

// ReleaseCreativeQueryVo 释放CreativeQueryVo
func ReleaseCreativeQueryVo(v *CreativeQueryVo) {
	v.AdgroupIdList = v.AdgroupIdList[:0]
	v.CampaignIdList = v.CampaignIdList[:0]
	v.Offset = 0
	v.PageSize = 0
	v.RptQuery = nil
	v.CreativeId = 0
	poolCreativeQueryVo.Put(v)
}
