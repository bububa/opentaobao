package logistic

// AddressReachableTopResult 
type AddressReachableTopResult struct {
    // 筛单结果l列表
    ReachableResultList   []AddressReachableResult `json:"reachable_result_list,omitempty" xml:"reachable_result_list>address_reachable_result,omitempty"`
}
