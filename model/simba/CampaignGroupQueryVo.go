package simba

import (
	"sync"
)

// CampaignGroupQueryVo 结构体
type CampaignGroupQueryVo struct {
	// 业务身份,枚举信息同&#39;account.get.can.use.bizcode&#39;服务返回结果
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 计划组名称
	CampaignGroupName string `json:"campaign_group_name,omitempty" xml:"campaign_group_name,omitempty"`
}

var poolCampaignGroupQueryVo = sync.Pool{
	New: func() any {
		return new(CampaignGroupQueryVo)
	},
}

// GetCampaignGroupQueryVo() 从对象池中获取CampaignGroupQueryVo
func GetCampaignGroupQueryVo() *CampaignGroupQueryVo {
	return poolCampaignGroupQueryVo.Get().(*CampaignGroupQueryVo)
}

// ReleaseCampaignGroupQueryVo 释放CampaignGroupQueryVo
func ReleaseCampaignGroupQueryVo(v *CampaignGroupQueryVo) {
	v.BizCode = ""
	v.CampaignGroupName = ""
	poolCampaignGroupQueryVo.Put(v)
}
