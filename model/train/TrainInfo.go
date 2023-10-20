package train

import (
	"sync"
)

// TrainInfo 结构体
type TrainInfo struct {
	// 到达站
	TrainTo string `json:"train_to,omitempty" xml:"train_to,omitempty"`
	// 出发站
	TrainFrom string `json:"train_from,omitempty" xml:"train_from,omitempty"`
	// 票号
	SequenceNo string `json:"sequence_no,omitempty" xml:"sequence_no,omitempty"`
	// 出发时间
	TrainDate string `json:"train_date,omitempty" xml:"train_date,omitempty"`
	// 车次号
	TrainNo string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// 坐席信息
	ChooseSeat string `json:"choose_seat,omitempty" xml:"choose_seat,omitempty"`
}

var poolTrainInfo = sync.Pool{
	New: func() any {
		return new(TrainInfo)
	},
}

// GetTrainInfo() 从对象池中获取TrainInfo
func GetTrainInfo() *TrainInfo {
	return poolTrainInfo.Get().(*TrainInfo)
}

// ReleaseTrainInfo 释放TrainInfo
func ReleaseTrainInfo(v *TrainInfo) {
	v.TrainTo = ""
	v.TrainFrom = ""
	v.SequenceNo = ""
	v.TrainDate = ""
	v.TrainNo = ""
	v.ChooseSeat = ""
	poolTrainInfo.Put(v)
}
