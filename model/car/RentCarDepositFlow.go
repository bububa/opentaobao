package car

import (
	"sync"
)

// RentCarDepositFlow 结构体
type RentCarDepositFlow struct {
	// 主标题
	MainTitle string `json:"main_title,omitempty" xml:"main_title,omitempty"`
	// 发生时间
	OccurTime string `json:"occur_time,omitempty" xml:"occur_time,omitempty"`
	// 状态图标
	StatusIcon string `json:"status_icon,omitempty" xml:"status_icon,omitempty"`
	// 子标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolRentCarDepositFlow = sync.Pool{
	New: func() any {
		return new(RentCarDepositFlow)
	},
}

// GetRentCarDepositFlow() 从对象池中获取RentCarDepositFlow
func GetRentCarDepositFlow() *RentCarDepositFlow {
	return poolRentCarDepositFlow.Get().(*RentCarDepositFlow)
}

// ReleaseRentCarDepositFlow 释放RentCarDepositFlow
func ReleaseRentCarDepositFlow(v *RentCarDepositFlow) {
	v.MainTitle = ""
	v.OccurTime = ""
	v.StatusIcon = ""
	v.SubTitle = ""
	v.Status = 0
	poolRentCarDepositFlow.Put(v)
}
