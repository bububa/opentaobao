package xhotel

// HbsResult 
type HbsResult struct {
    // 错误信息
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 飞猪平台订单号
    Module   int64 `json:"module,omitempty" xml:"module,omitempty"`
    // 结果码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 给开发用的错误信息
    ResultMsg4Dev   string `json:"result_msg4_dev,omitempty" xml:"result_msg4_dev,omitempty"`
    // 订单是否创建成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
