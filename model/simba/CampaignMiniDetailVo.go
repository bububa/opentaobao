package simba

import (
	"sync"
)

// CampaignMiniDetailVo 结构体
type CampaignMiniDetailVo struct {
	// 橱窗宝贝id列表
	PrivateMiniItemIdList []int64 `json:"private_mini_item_id_list,omitempty" xml:"private_mini_item_id_list>int64,omitempty"`
	// 橱窗宝贝里面置顶的id列表
	PrivateMiniTopItemIdList []int64 `json:"private_mini_top_item_id_list,omitempty" xml:"private_mini_top_item_id_list>int64,omitempty"`
	// 橱窗主题,1:默认主题,7:品牌定制主题,4:店铺主题,5:拉新主题,2:新品主题
	MiniDetailTheme string `json:"mini_detail_theme,omitempty" xml:"mini_detail_theme,omitempty"`
	// 手动选宝贝还是智能,0:黑名单,1:白名单
	PrivateMiniType int64 `json:"private_mini_type,omitempty" xml:"private_mini_type,omitempty"`
}

var poolCampaignMiniDetailVo = sync.Pool{
	New: func() any {
		return new(CampaignMiniDetailVo)
	},
}

// GetCampaignMiniDetailVo() 从对象池中获取CampaignMiniDetailVo
func GetCampaignMiniDetailVo() *CampaignMiniDetailVo {
	return poolCampaignMiniDetailVo.Get().(*CampaignMiniDetailVo)
}

// ReleaseCampaignMiniDetailVo 释放CampaignMiniDetailVo
func ReleaseCampaignMiniDetailVo(v *CampaignMiniDetailVo) {
	v.PrivateMiniItemIdList = v.PrivateMiniItemIdList[:0]
	v.PrivateMiniTopItemIdList = v.PrivateMiniTopItemIdList[:0]
	v.MiniDetailTheme = ""
	v.PrivateMiniType = 0
	poolCampaignMiniDetailVo.Put(v)
}
