package film

// ResultListModel 
/* model for simplify = false
type ResultListModel struct {

    // 错误码
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // 返回值
    
    AccountList  struct {
        TaobaoFilmAccountPhoneQueryModel  []TaobaoFilmAccountPhoneQueryModel `json:"taobao_film_account_phone_query_model,omitempty"`
    } `json:"account_list,omitempty"`
    

    // 请求ID
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 错误描述
    
    ReturnMessage   string `json:"return_message,omitempty"`
    

    // 忽略
    
    ReturnUrl   string `json:"return_url,omitempty"`
    

    // returnValue
    
    ReturnValues  struct {
        LotteryDraws  []LotteryDraws `json:"lottery_draws,omitempty"`
    } `json:"return_values,omitempty"`
    

    // returnErrorOper
    
    ReturnErrorOper   string `json:"return_error_oper,omitempty"`
    

    // returnErrorSolution
    
    ReturnErrorSolution   string `json:"return_error_solution,omitempty"`
    

    // returnErrorStackTrace
    
    ReturnErrorStackTrace   string `json:"return_error_stack_trace,omitempty"`
    

}
*/

// ResultListModel 
type ResultListModel struct {

    // 错误码
    ReturnCode   string `json:"return_code,omitempty"`

    // 返回值
    AccountList   []TaobaoFilmAccountPhoneQueryModel `json:"account_list,omitempty"`

    // 请求ID
    RequestId   string `json:"request_id,omitempty"`

    // 错误描述
    ReturnMessage   string `json:"return_message,omitempty"`

    // 忽略
    ReturnUrl   string `json:"return_url,omitempty"`

    // returnValue
    ReturnValues   []LotteryDraws `json:"return_values,omitempty"`

    // returnErrorOper
    ReturnErrorOper   string `json:"return_error_oper,omitempty"`

    // returnErrorSolution
    ReturnErrorSolution   string `json:"return_error_solution,omitempty"`

    // returnErrorStackTrace
    ReturnErrorStackTrace   string `json:"return_error_stack_trace,omitempty"`

}
