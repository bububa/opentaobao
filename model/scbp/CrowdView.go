package scbp

import (
	"sync"
)

// CrowdView 结构体
type CrowdView struct {
	// 人群list
	CrowdList []CrowdInfo `json:"crowd_list,omitempty" xml:"crowd_list>crowd_info,omitempty"`
	// 人群类型(店铺老客、优选人群)
	CrowdType string `json:"crowd_type,omitempty" xml:"crowd_type,omitempty"`
}

var poolCrowdView = sync.Pool{
	New: func() any {
		return new(CrowdView)
	},
}

// GetCrowdView() 从对象池中获取CrowdView
func GetCrowdView() *CrowdView {
	return poolCrowdView.Get().(*CrowdView)
}

// ReleaseCrowdView 释放CrowdView
func ReleaseCrowdView(v *CrowdView) {
	v.CrowdList = v.CrowdList[:0]
	v.CrowdType = ""
	poolCrowdView.Put(v)
}
