package tbk

import (
	"sync"
)

// SpCampaign 结构体
type SpCampaign struct {
	// 定向计划活动ID
	SpCid string `json:"sp_cid,omitempty" xml:"sp_cid,omitempty"`
	// 定向计划名称
	SpName string `json:"sp_name,omitempty" xml:"sp_name,omitempty"`
	// 定向佣金率
	SpRate string `json:"sp_rate,omitempty" xml:"sp_rate,omitempty"`
	// 定向是否锁佣，0=不锁佣 1=锁佣
	SpLockStatus string `json:"sp_lock_status,omitempty" xml:"sp_lock_status,omitempty"`
	// 定向计划申请链接
	SpApplyLink string `json:"sp_apply_link,omitempty" xml:"sp_apply_link,omitempty"`
	// 定向计划是否可用 1-可用 0-不可用
	SpStatus string `json:"sp_status,omitempty" xml:"sp_status,omitempty"`
}

var poolSpCampaign = sync.Pool{
	New: func() any {
		return new(SpCampaign)
	},
}

// GetSpCampaign() 从对象池中获取SpCampaign
func GetSpCampaign() *SpCampaign {
	return poolSpCampaign.Get().(*SpCampaign)
}

// ReleaseSpCampaign 释放SpCampaign
func ReleaseSpCampaign(v *SpCampaign) {
	v.SpCid = ""
	v.SpName = ""
	v.SpRate = ""
	v.SpLockStatus = ""
	v.SpApplyLink = ""
	v.SpStatus = ""
	poolSpCampaign.Put(v)
}
