package btrip

import (
	"sync"
)

// TrainOrderInfo 结构体
type TrainOrderInfo struct {
	// 车次信息
	TrainInfoList []TrainInfoV2 `json:"train_info_list,omitempty" xml:"train_info_list>train_info_v2,omitempty"`
	// 中转信息
	TrainTransferInfo *TrainTransferInfo `json:"train_transfer_info,omitempty" xml:"train_transfer_info,omitempty"`
}

var poolTrainOrderInfo = sync.Pool{
	New: func() any {
		return new(TrainOrderInfo)
	},
}

// GetTrainOrderInfo() 从对象池中获取TrainOrderInfo
func GetTrainOrderInfo() *TrainOrderInfo {
	return poolTrainOrderInfo.Get().(*TrainOrderInfo)
}

// ReleaseTrainOrderInfo 释放TrainOrderInfo
func ReleaseTrainOrderInfo(v *TrainOrderInfo) {
	v.TrainInfoList = v.TrainInfoList[:0]
	v.TrainTransferInfo = nil
	poolTrainOrderInfo.Put(v)
}
