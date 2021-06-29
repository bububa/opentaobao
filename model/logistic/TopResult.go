package logistic

// TopResult 
type TopResult struct {
    // 发货完成后的物流单号
    LgOrderCode   string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
    // 错误信息
    SubErrorMsg   string `json:"sub_error_msg,omitempty" xml:"sub_error_msg,omitempty"`
    // 错误code
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误code
    SubErrorCode   string `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 资源列表
    Resources   []ThreePlConsignResourceDTO `json:"resources,omitempty" xml:"resources>three_pl_consign_resource_dto,omitempty"`
}
