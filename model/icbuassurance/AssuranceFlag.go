package icbuassurance

import (
	"sync"
)

// AssuranceFlag 结构体
type AssuranceFlag struct {
	// flagList
	FlagList []string `json:"flag_list,omitempty" xml:"flag_list>string,omitempty"`
	// guideURL
	GuideURL string `json:"guide_u_r_l,omitempty" xml:"guide_u_r_l,omitempty"`
	// pauseStatus
	PauseStatus bool `json:"pause_status,omitempty" xml:"pause_status,omitempty"`
}

var poolAssuranceFlag = sync.Pool{
	New: func() any {
		return new(AssuranceFlag)
	},
}

// GetAssuranceFlag() 从对象池中获取AssuranceFlag
func GetAssuranceFlag() *AssuranceFlag {
	return poolAssuranceFlag.Get().(*AssuranceFlag)
}

// ReleaseAssuranceFlag 释放AssuranceFlag
func ReleaseAssuranceFlag(v *AssuranceFlag) {
	v.FlagList = v.FlagList[:0]
	v.GuideURL = ""
	v.PauseStatus = false
	poolAssuranceFlag.Put(v)
}
