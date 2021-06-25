package refund

// TaobaoRdcAligeniusSendgoodsCancelResult 
type TaobaoRdcAligeniusSendgoodsCancelResult struct {

    // resultData
    ResultData   *Resultdata `json:"result_data,omitempty"`

    // 异常信息
    ErrorInfo   string `json:"error_info,omitempty"`

    // 异常编码
    ErrorCode   string `json:"error_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
