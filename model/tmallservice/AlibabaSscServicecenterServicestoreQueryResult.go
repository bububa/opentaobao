package tmallservice

// AlibabaSscServicecenterServicestoreQueryResult 
/* model for simplify = false
type AlibabaSscServicecenterServicestoreQueryResult struct {

    // 明细条目执行结果对象
    
    ResultDataList  struct {
        ResultData  []ResultData `json:"result_data,omitempty"`
    } `json:"result_data_list,omitempty"`
    

}
*/

// AlibabaSscServicecenterServicestoreQueryResult 
type AlibabaSscServicecenterServicestoreQueryResult struct {

    // 明细条目执行结果对象
    ResultDataList   []ResultData `json:"result_data_list,omitempty"`

}
