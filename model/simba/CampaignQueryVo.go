package simba

import (
	"sync"
)

// CampaignQueryVo 结构体
type CampaignQueryVo struct {
	// 筛选项-主体类型
	StatusList []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
	// 筛选项-主体类型
	PromotionTypeList []string `json:"promotion_type_list,omitempty" xml:"promotion_type_list>string,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolCampaignQueryVo = sync.Pool{
	New: func() any {
		return new(CampaignQueryVo)
	},
}

// GetCampaignQueryVo() 从对象池中获取CampaignQueryVo
func GetCampaignQueryVo() *CampaignQueryVo {
	return poolCampaignQueryVo.Get().(*CampaignQueryVo)
}

// ReleaseCampaignQueryVo 释放CampaignQueryVo
func ReleaseCampaignQueryVo(v *CampaignQueryVo) {
	v.StatusList = v.StatusList[:0]
	v.PromotionTypeList = v.PromotionTypeList[:0]
	v.CampaignId = 0
	v.Offset = 0
	v.PageSize = 0
	poolCampaignQueryVo.Put(v)
}
