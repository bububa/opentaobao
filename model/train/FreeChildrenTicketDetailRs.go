package train

import (
	"sync"
)

// FreeChildrenTicketDetailRs 结构体
type FreeChildrenTicketDetailRs struct {
	// 唯一标识
	ApplyNo string `json:"apply_no,omitempty" xml:"apply_no,omitempty"`
	// 超时时间
	Timeout string `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// 火车信息
	TrainInfo *TrainInfo `json:"train_info,omitempty" xml:"train_info,omitempty"`
	// 成人信息
	PassengerInfo *PassengerInfo `json:"passenger_info,omitempty" xml:"passenger_info,omitempty"`
	// 儿童信息
	FreeChildrenPassengerInfo *PassengerInfo `json:"free_children_passenger_info,omitempty" xml:"free_children_passenger_info,omitempty"`
	// 操作类型
	OperatorType int64 `json:"operator_type,omitempty" xml:"operator_type,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolFreeChildrenTicketDetailRs = sync.Pool{
	New: func() any {
		return new(FreeChildrenTicketDetailRs)
	},
}

// GetFreeChildrenTicketDetailRs() 从对象池中获取FreeChildrenTicketDetailRs
func GetFreeChildrenTicketDetailRs() *FreeChildrenTicketDetailRs {
	return poolFreeChildrenTicketDetailRs.Get().(*FreeChildrenTicketDetailRs)
}

// ReleaseFreeChildrenTicketDetailRs 释放FreeChildrenTicketDetailRs
func ReleaseFreeChildrenTicketDetailRs(v *FreeChildrenTicketDetailRs) {
	v.ApplyNo = ""
	v.Timeout = ""
	v.TrainInfo = nil
	v.PassengerInfo = nil
	v.FreeChildrenPassengerInfo = nil
	v.OperatorType = 0
	v.Status = 0
	poolFreeChildrenTicketDetailRs.Put(v)
}
