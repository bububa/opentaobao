package scbp

import (
	"sync"
)

// CampaignQueryDto 结构体
type CampaignQueryDto struct {
	// 计划类型列表
	TypeList []string `json:"type_list,omitempty" xml:"type_list>string,omitempty"`
	// 计划id集合
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 计划标题，配合exactMatch使用
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 子类型
	SubType string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
	// 推广状态
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 类目id
	CateId int64 `json:"cate_id,omitempty" xml:"cate_id,omitempty"`
	// 当前页数
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 标题是否精确匹配
	ExactMatch bool `json:"exact_match,omitempty" xml:"exact_match,omitempty"`
}

var poolCampaignQueryDto = sync.Pool{
	New: func() any {
		return new(CampaignQueryDto)
	},
}

// GetCampaignQueryDto() 从对象池中获取CampaignQueryDto
func GetCampaignQueryDto() *CampaignQueryDto {
	return poolCampaignQueryDto.Get().(*CampaignQueryDto)
}

// ReleaseCampaignQueryDto 释放CampaignQueryDto
func ReleaseCampaignQueryDto(v *CampaignQueryDto) {
	v.TypeList = v.TypeList[:0]
	v.CampaignIdList = v.CampaignIdList[:0]
	v.Title = ""
	v.SubType = ""
	v.OnlineStatus = 0
	v.CateId = 0
	v.Page = 0
	v.Size = 0
	v.ExactMatch = false
	poolCampaignQueryDto.Put(v)
}
