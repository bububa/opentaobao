package logistic

// AddressReachableTopResult 
/* model for simplify = false
type AddressReachableTopResult struct {

    // 筛单结果l列表
    
    ReachableResultList  struct {
        AddressReachableResult  []AddressReachableResult `json:"address_reachable_result,omitempty"`
    } `json:"reachable_result_list,omitempty"`
    

}
*/

// AddressReachableTopResult 
type AddressReachableTopResult struct {

    // 筛单结果l列表
    ReachableResultList   []AddressReachableResult `json:"reachable_result_list,omitempty"`

}
