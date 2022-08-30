package wlbimports

// TransferStoreInfoResponse 结构体
type TransferStoreInfoResponse struct {
	// 仓库揽收信息
	PickUpResponse *PickupInfo `json:"pick_up_response,omitempty" xml:"pick_up_response,omitempty"`
	// 集货仓信息
	StoreCenterResponse *TransferStoreCenterResponse `json:"store_center_response,omitempty" xml:"store_center_response,omitempty"`
}
