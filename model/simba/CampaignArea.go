package simba

import (
	"sync"
)

// CampaignArea 结构体
type CampaignArea struct {
	// 主人昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID，如果已经包含省、自治区id，请不要再包括省内市的id；
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 最后修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 推广计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolCampaignArea = sync.Pool{
	New: func() any {
		return new(CampaignArea)
	},
}

// GetCampaignArea() 从对象池中获取CampaignArea
func GetCampaignArea() *CampaignArea {
	return poolCampaignArea.Get().(*CampaignArea)
}

// ReleaseCampaignArea 释放CampaignArea
func ReleaseCampaignArea(v *CampaignArea) {
	v.Nick = ""
	v.Area = ""
	v.CreateTime = ""
	v.ModifiedTime = ""
	v.CampaignId = 0
	poolCampaignArea.Put(v)
}
