package btrip

// BtripApplyResult 
type BtripApplyResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 结果对象
    Module   *OpenApiNewApplyRs `json:"module,omitempty" xml:"module,omitempty"`
    // 错误信息
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 结果码
    ErrCode   int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 结果描述
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 审批单列表
    ApplyList   []OpenApplyRs `json:"apply_list,omitempty" xml:"apply_list>open_apply_rs,omitempty"`
    // 总数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
}
