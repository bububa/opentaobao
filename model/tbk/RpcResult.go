package tbk

// RpcResult 
type RpcResult struct {

    // 结果
    Data   *PageResult `json:"data,omitempty"`

    // 业务出错的描述
    BizErrorDesc   string `json:"biz_error_desc,omitempty"`

    // 业务出错的状态码
    BizErrorCode   int64 `json:"biz_error_code,omitempty"`

    // 执行结果
    ResultMsg   string `json:"result_msg,omitempty"`

    // 执行结果状态码
    ResultCode   int64 `json:"result_code,omitempty"`

}
