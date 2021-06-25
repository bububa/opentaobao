package bus

// TaobaoBusLastplaceGetResult 
type TaobaoBusLastplaceGetResult struct {

    // errCode
    ErrCode   string `json:"err_code,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // 目的地数据JSON,具体参考文档说明
    Module   string `json:"module,omitempty"`

}
