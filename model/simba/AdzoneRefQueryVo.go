package simba

import (
	"sync"
)

// AdzoneRefQueryVo 结构体
type AdzoneRefQueryVo struct {
	// 计划id集合
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolAdzoneRefQueryVo = sync.Pool{
	New: func() any {
		return new(AdzoneRefQueryVo)
	},
}

// GetAdzoneRefQueryVo() 从对象池中获取AdzoneRefQueryVo
func GetAdzoneRefQueryVo() *AdzoneRefQueryVo {
	return poolAdzoneRefQueryVo.Get().(*AdzoneRefQueryVo)
}

// ReleaseAdzoneRefQueryVo 释放AdzoneRefQueryVo
func ReleaseAdzoneRefQueryVo(v *AdzoneRefQueryVo) {
	v.CampaignIdList = v.CampaignIdList[:0]
	v.Offset = 0
	v.PageSize = 0
	poolAdzoneRefQueryVo.Put(v)
}
