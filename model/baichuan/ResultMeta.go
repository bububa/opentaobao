package baichuan

// ResultMeta 
type ResultMeta struct {

    // 返回码
    Code   int64 `json:"code,omitempty"`

    // 返回码对应的文案
    Msg   string `json:"msg,omitempty"`

    // 返回的详细内容
    Data   *ResultData `json:"data,omitempty"`

}
