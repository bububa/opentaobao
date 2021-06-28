package refund

// TaobaoRdcAligeniusIdentificationCaseUpdateResult 
/* model for simplify = false
type TaobaoRdcAligeniusIdentificationCaseUpdateResult struct {

    // resultData
    
    ResultData  *struct {
        Resultdata  *Resultdata `json:"resultdata,omitempty"`
    } `json:"result_data,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorInfo   string `json:"error_info,omitempty"`
    

}
*/

// TaobaoRdcAligeniusIdentificationCaseUpdateResult 
type TaobaoRdcAligeniusIdentificationCaseUpdateResult struct {

    // resultData
    ResultData   *Resultdata `json:"result_data,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorInfo   string `json:"error_info,omitempty"`

}
