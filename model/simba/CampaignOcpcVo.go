package simba

import (
	"sync"
)

// CampaignOcpcVo 结构体
type CampaignOcpcVo struct {
	// OCPC是否开启，0:否，1:是
	EnableOcpc string `json:"enable_ocpc,omitempty" xml:"enable_ocpc,omitempty"`
	// OCPC溢价比例
	OcpcRatio int64 `json:"ocpc_ratio,omitempty" xml:"ocpc_ratio,omitempty"`
}

var poolCampaignOcpcVo = sync.Pool{
	New: func() any {
		return new(CampaignOcpcVo)
	},
}

// GetCampaignOcpcVo() 从对象池中获取CampaignOcpcVo
func GetCampaignOcpcVo() *CampaignOcpcVo {
	return poolCampaignOcpcVo.Get().(*CampaignOcpcVo)
}

// ReleaseCampaignOcpcVo 释放CampaignOcpcVo
func ReleaseCampaignOcpcVo(v *CampaignOcpcVo) {
	v.EnableOcpc = ""
	v.OcpcRatio = 0
	poolCampaignOcpcVo.Put(v)
}
