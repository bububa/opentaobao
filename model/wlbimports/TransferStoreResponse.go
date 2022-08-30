package wlbimports

// TransferStoreResponse 结构体
type TransferStoreResponse struct {
	// 集货仓信息列表
	TransferStoreInfoResponseList []TransferStoreInfoResponse `json:"transfer_store_info_response_list,omitempty" xml:"transfer_store_info_response_list>transfer_store_info_response,omitempty"`
}
