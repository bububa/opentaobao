package mos

// MultiResult 
/* model for simplify = false
type MultiResult struct {

    // total
    
    ResultTotal   int64 `json:"result_total,omitempty"`
    

    // errMessage
    
    ResultMessage   string `json:"result_message,omitempty"`
    

    // data
    
    ResultDatas  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"result_datas,omitempty"`
    

    // errCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // success
    
    ResultSuccess   bool `json:"result_success,omitempty"`
    

}
*/

// MultiResult 
type MultiResult struct {

    // total
    ResultTotal   int64 `json:"result_total,omitempty"`

    // errMessage
    ResultMessage   string `json:"result_message,omitempty"`

    // data
    ResultDatas   []string `json:"result_datas,omitempty"`

    // errCode
    ResultCode   string `json:"result_code,omitempty"`

    // success
    ResultSuccess   bool `json:"result_success,omitempty"`

}
