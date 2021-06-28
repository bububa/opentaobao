package refund

// TaobaoRdcAligeniusSendgoodsCancelResult 
/* model for simplify = false
type TaobaoRdcAligeniusSendgoodsCancelResult struct {

    // resultData
    
    ResultData  *struct {
        Resultdata  *Resultdata `json:"resultdata,omitempty"`
    } `json:"result_data,omitempty"`
    

    // 异常信息
    
    ErrorInfo   string `json:"error_info,omitempty"`
    

    // 异常编码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

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
