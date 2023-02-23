package btrip

// TrainOrderInfo 结构体
type TrainOrderInfo struct {
	// 车次信息
	TrainInfoList []TrainInfoV2 `json:"train_info_list,omitempty" xml:"train_info_list>train_info_v2,omitempty"`
	// 中转信息
	TrainTransferInfo *TrainTransferInfo `json:"train_transfer_info,omitempty" xml:"train_transfer_info,omitempty"`
}
