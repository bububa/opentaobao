package simba

import (
	"sync"
)

// CampaignLaunchTimeVo 结构体
type CampaignLaunchTimeVo struct {
	// 计划开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 计划结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 是否长期投放,true:是,false:否
	LaunchForever bool `json:"launch_forever,omitempty" xml:"launch_forever,omitempty"`
}

var poolCampaignLaunchTimeVo = sync.Pool{
	New: func() any {
		return new(CampaignLaunchTimeVo)
	},
}

// GetCampaignLaunchTimeVo() 从对象池中获取CampaignLaunchTimeVo
func GetCampaignLaunchTimeVo() *CampaignLaunchTimeVo {
	return poolCampaignLaunchTimeVo.Get().(*CampaignLaunchTimeVo)
}

// ReleaseCampaignLaunchTimeVo 释放CampaignLaunchTimeVo
func ReleaseCampaignLaunchTimeVo(v *CampaignLaunchTimeVo) {
	v.StartTime = ""
	v.EndTime = ""
	v.LaunchForever = false
	poolCampaignLaunchTimeVo.Put(v)
}
