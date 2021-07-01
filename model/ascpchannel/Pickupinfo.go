package ascpchannel

// Pickupinfo 结构体
type Pickupinfo struct {
	// 集货是否完成 Y 是,N 否
	IsCompleted string `json:"is_completed,omitempty" xml:"is_completed,omitempty"`
	// 集货是否取消 Y 是,N 否
	IsCanceled string `json:"is_canceled,omitempty" xml:"is_canceled,omitempty"`
	// 集货发货单号列表（用,分隔）
	PickUpNos string `json:"pick_up_nos,omitempty" xml:"pick_up_nos,omitempty"`
}
