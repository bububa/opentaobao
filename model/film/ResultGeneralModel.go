package film

// ResultGeneralModel 
/* model for simplify = false
type ResultGeneralModel struct {

    // 调用失败描述信息
    
    ReturnMessage   string `json:"return_message,omitempty"`
    

    // 0代表接口调用成功，其他值表示调用失败，错误信息可详见接口文档
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // returnValue
    
    ReturnValue  *struct {
        TopRefundOrderStatus  *TopRefundOrderStatus `json:"top_refund_order_status,omitempty"`
    } `json:"return_value,omitempty"`
    

    // returnUrl
    
    ReturnUrl   string `json:"return_url,omitempty"`
    

    // requestId
    
    RequestId   string `json:"request_id,omitempty"`
    

    // returnErrorOper
    
    ReturnErrorOper   string `json:"return_error_oper,omitempty"`
    

    // returnErrorSolution
    
    ReturnErrorSolution   string `json:"return_error_solution,omitempty"`
    

    // returnErrorStackTrace
    
    ReturnErrorStackTrace   string `json:"return_error_stack_trace,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// ResultGeneralModel 
type ResultGeneralModel struct {

    // 调用失败描述信息
    ReturnMessage   string `json:"return_message,omitempty"`

    // 0代表接口调用成功，其他值表示调用失败，错误信息可详见接口文档
    ReturnCode   string `json:"return_code,omitempty"`

    // returnValue
    ReturnValue   *TopRefundOrderStatus `json:"return_value,omitempty"`

    // returnUrl
    ReturnUrl   string `json:"return_url,omitempty"`

    // requestId
    RequestId   string `json:"request_id,omitempty"`

    // returnErrorOper
    ReturnErrorOper   string `json:"return_error_oper,omitempty"`

    // returnErrorSolution
    ReturnErrorSolution   string `json:"return_error_solution,omitempty"`

    // returnErrorStackTrace
    ReturnErrorStackTrace   string `json:"return_error_stack_trace,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
