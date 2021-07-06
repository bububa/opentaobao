package logistic

// RoutingReachableBatchRequestDto 结构体
type RoutingReachableBatchRequestDto struct {
	// 收发地址和服务列表
	AddressAndServiceList []ReachableAddressAndServiceDto `json:"address_and_service_list,omitempty" xml:"address_and_service_list>reachable_address_and_service_dto,omitempty"`
	// 快递公司code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 发件揽收网点
	SendBranchCode string `json:"send_branch_code,omitempty" xml:"send_branch_code,omitempty"`
}
