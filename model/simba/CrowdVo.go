package simba

import (
	"sync"
)

// CrowdVo 结构体
type CrowdVo struct {
	// 种子人群关联信息
	ExtendSeedCrowdList []ExtendSeedCrowdRefVo `json:"extend_seed_crowd_list,omitempty" xml:"extend_seed_crowd_list>extend_seed_crowd_ref_vo,omitempty"`
	// 子人群信息
	SubCrowdList []SubCrowdVo `json:"sub_crowd_list,omitempty" xml:"sub_crowd_list>sub_crowd_vo,omitempty"`
	// 人群名称
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
	// 人群值
	CrowdValue string `json:"crowd_value,omitempty" xml:"crowd_value,omitempty"`
	// 扩展倍数。低中高对应不同的倍数，由产品定义具体值，未来业务可支持滑动条。注意：倍数是指最终人数/种子人数，比如种子人群100万，扩展比种子多了50万，也就是总共150万，这时候扩展倍数是1.5而不是0.5
	LookalikeMultiple string `json:"lookalike_multiple,omitempty" xml:"lookalike_multiple,omitempty"`
	// 人群主键id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 定向类型
	TargetType int64 `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 定向标签
	Label *LabelVo `json:"label,omitempty" xml:"label,omitempty"`
}

var poolCrowdVo = sync.Pool{
	New: func() any {
		return new(CrowdVo)
	},
}

// GetCrowdVo() 从对象池中获取CrowdVo
func GetCrowdVo() *CrowdVo {
	return poolCrowdVo.Get().(*CrowdVo)
}

// ReleaseCrowdVo 释放CrowdVo
func ReleaseCrowdVo(v *CrowdVo) {
	v.ExtendSeedCrowdList = v.ExtendSeedCrowdList[:0]
	v.SubCrowdList = v.SubCrowdList[:0]
	v.CrowdName = ""
	v.CrowdValue = ""
	v.LookalikeMultiple = ""
	v.CrowdId = 0
	v.TargetType = 0
	v.Label = nil
	poolCrowdVo.Put(v)
}
