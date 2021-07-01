package flight

// DelAuxProductsRS 
type DelAuxProductsRS struct {
    // 操作日志id，商家可通过该id在后台查看本次操作的具体结果
    TracerId   string `json:"tracer_id,omitempty" xml:"tracer_id,omitempty"`
    // 备注
    ApiErrorMsg   string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
    // 是否操作成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
