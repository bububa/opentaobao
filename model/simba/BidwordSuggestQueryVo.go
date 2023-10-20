package simba

import (
	"sync"
)

// BidwordSuggestQueryVo 结构体
type BidwordSuggestQueryVo struct {
	// 人群id
	CrowdIdList []int64 `json:"crowd_id_list,omitempty" xml:"crowd_id_list>int64,omitempty"`
	// 类型,kr_overall:综合推荐,kr_flow:精准引流,kr_category:类目优质词,kr_industry:行业机会词,kr_mta:助攻词,kr_search:搜索词联想,kr_byword:以词推词,kr_bycrowd:以人群词,new_default:新建流程默认推词
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 宝贝id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// 单元id,单元已经存在场景必填
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 计划id,计划已经存在场景必填
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolBidwordSuggestQueryVo = sync.Pool{
	New: func() any {
		return new(BidwordSuggestQueryVo)
	},
}

// GetBidwordSuggestQueryVo() 从对象池中获取BidwordSuggestQueryVo
func GetBidwordSuggestQueryVo() *BidwordSuggestQueryVo {
	return poolBidwordSuggestQueryVo.Get().(*BidwordSuggestQueryVo)
}

// ReleaseBidwordSuggestQueryVo 释放BidwordSuggestQueryVo
func ReleaseBidwordSuggestQueryVo(v *BidwordSuggestQueryVo) {
	v.CrowdIdList = v.CrowdIdList[:0]
	v.Type = ""
	v.MaterialId = 0
	v.AdgroupId = 0
	v.CampaignId = 0
	poolBidwordSuggestQueryVo.Put(v)
}
