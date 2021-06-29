package logistic

// ReachableServiceWaybillResponseDTO 
type ReachableServiceWaybillResponseDTO struct {
    // 单个结果是否异常
    ErrorInfo   *ErrorInfo `json:"error_info,omitempty" xml:"error_info,omitempty"`
    // 单个结果是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 可达信息
    Module   *ReachableDTO `json:"module,omitempty" xml:"module,omitempty"`
    // 与入参地址列表中单项objectId对应
    ObjectId   string `json:"object_id,omitempty" xml:"object_id,omitempty"`
}
