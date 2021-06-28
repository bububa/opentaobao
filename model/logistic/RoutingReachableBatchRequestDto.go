package logistic

// RoutingReachableBatchRequestDto 
/* model for simplify = false
type RoutingReachableBatchRequestDto struct {

    // 快递公司code
    
    CpCode   string `json:"cp_code,omitempty"`
    

    // 收发地址和服务列表
    
    AddressAndServiceList  struct {
        ReachableAddressAndServiceDto  []ReachableAddressAndServiceDto `json:"reachable_address_and_service_dto,omitempty"`
    } `json:"address_and_service_list,omitempty"`
    

    // 发件揽收网点
    
    SendBranchCode   string `json:"send_branch_code,omitempty"`
    

}
*/

// RoutingReachableBatchRequestDto 
type RoutingReachableBatchRequestDto struct {

    // 快递公司code
    CpCode   string `json:"cp_code,omitempty"`

    // 收发地址和服务列表
    AddressAndServiceList   []ReachableAddressAndServiceDto `json:"address_and_service_list,omitempty"`

    // 发件揽收网点
    SendBranchCode   string `json:"send_branch_code,omitempty"`

}
