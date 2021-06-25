package tmallhk

// DataResult 
type DataResult struct {

    // 参数code
    Code   string `json:"code,omitempty"`

    // 参数msg
    Msg   string `json:"msg,omitempty"`

    // 是否正常
    Success   bool `json:"success,omitempty"`

    // obj
    Obj   *CCICCheckCodeDO `json:"obj,omitempty"`

}
