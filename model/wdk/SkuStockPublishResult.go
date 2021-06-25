package wdk

// SkuStockPublishResult 
type SkuStockPublishResult struct {

    // 具体的错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // bill_no对应的操作结果
    Result   bool `json:"result,omitempty"`

    // 入参中对应的幂等单据号
    BillNo   string `json:"bill_no,omitempty"`

    // 具体的更新失败原因
    ErrMsg   string `json:"err_msg,omitempty"`

}
