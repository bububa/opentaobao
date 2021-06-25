package mos

// AlibabaMosBunkBunkinfoQuerybunkResultDo 
type AlibabaMosBunkBunkinfoQuerybunkResultDo struct {

    // 全链路追踪id
    TraceId   string `json:"trace_id,omitempty"`

    // 总量
    Total   int64 `json:"total,omitempty"`

    // 返回数据
    DataList   []BunkSimpleDto `json:"data_list,omitempty"`

    // 调用是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrCode   int64 `json:"err_code,omitempty"`

    // 其他
    Extra   string `json:"extra,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 结果码
    ResultCode   string `json:"result_code,omitempty"`

    // 结果标题
    Titles   []String `json:"titles,omitempty"`

}