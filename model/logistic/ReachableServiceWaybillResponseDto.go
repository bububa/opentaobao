package logistic

// ReachableServiceWaybillResponseDto 
/* model for simplify = false
type ReachableServiceWaybillResponseDto struct {

    // 单个结果是否异常
    
    ErrorInfo  *struct {
        ErrorInfo  *ErrorInfo `json:"error_info,omitempty"`
    } `json:"error_info,omitempty"`
    

    // 单个结果是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 可达信息
    
    Module  *struct {
        ReachableDto  *ReachableDto `json:"reachable_dto,omitempty"`
    } `json:"module,omitempty"`
    

    // 与入参地址列表中单项objectId对应
    
    ObjectId   string `json:"object_id,omitempty"`
    

}
*/

// ReachableServiceWaybillResponseDto 
type ReachableServiceWaybillResponseDto struct {

    // 单个结果是否异常
    ErrorInfo   *ErrorInfo `json:"error_info,omitempty"`

    // 单个结果是否成功
    Success   bool `json:"success,omitempty"`

    // 可达信息
    Module   *ReachableDto `json:"module,omitempty"`

    // 与入参地址列表中单项objectId对应
    ObjectId   string `json:"object_id,omitempty"`

}
