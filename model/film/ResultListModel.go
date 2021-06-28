package film

// ResultListModel 
type ResultListModel struct {

    // 错误码
    
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`
    

    // 返回值
    
    AccountList   []TaobaoFilmAccountPhoneQueryModel `json:"account_list,omitempty" xml:"account_list,omitempty"`
    

    // 请求ID
    
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    

    // 错误描述
    
    ReturnMessage   string `json:"return_message,omitempty" xml:"return_message,omitempty"`
    

    // 忽略
    
    ReturnUrl   string `json:"return_url,omitempty" xml:"return_url,omitempty"`
    

    // returnValue
    
    ReturnValues   []LotteryDraws `json:"return_values,omitempty" xml:"return_values,omitempty"`
    

    // returnErrorOper
    
    ReturnErrorOper   string `json:"return_error_oper,omitempty" xml:"return_error_oper,omitempty"`
    

    // returnErrorSolution
    
    ReturnErrorSolution   string `json:"return_error_solution,omitempty" xml:"return_error_solution,omitempty"`
    

    // returnErrorStackTrace
    
    ReturnErrorStackTrace   string `json:"return_error_stack_trace,omitempty" xml:"return_error_stack_trace,omitempty"`
    

}
