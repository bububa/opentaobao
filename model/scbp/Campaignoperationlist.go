package scbp

import (
	"sync"
)

// Campaignoperationlist 结构体
type Campaignoperationlist struct {
	// 计划id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolCampaignoperationlist = sync.Pool{
	New: func() any {
		return new(Campaignoperationlist)
	},
}

// GetCampaignoperationlist() 从对象池中获取Campaignoperationlist
func GetCampaignoperationlist() *Campaignoperationlist {
	return poolCampaignoperationlist.Get().(*Campaignoperationlist)
}

// ReleaseCampaignoperationlist 释放Campaignoperationlist
func ReleaseCampaignoperationlist(v *Campaignoperationlist) {
	v.Id = 0
	poolCampaignoperationlist.Put(v)
}
