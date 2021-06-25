package iot

// BaseResult 
type BaseResult struct {

    // 信息
    Message   string `json:"message,omitempty"`

    // 响应code
    RetCode   int64 `json:"ret_code,omitempty"`

    // 返回结果
    RetValue   *BusinessRecipeOpenDto `json:"ret_value,omitempty"`

    // 追踪id
    TraceId   string `json:"trace_id,omitempty"`

}
