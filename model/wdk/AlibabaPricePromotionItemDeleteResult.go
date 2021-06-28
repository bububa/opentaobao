package wdk

// AlibabaPricePromotionItemDeleteResult 
/* model for simplify = false
type AlibabaPricePromotionItemDeleteResult struct {

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // data
    
    DataList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"data_list,omitempty"`
    

    // code
    
    ResultCode   int64 `json:"result_code,omitempty"`
    

    // totalRecord
    
    TotalRecord   int64 `json:"total_record,omitempty"`
    

    // msg
    
    Message   string `json:"message,omitempty"`
    

}
*/

// AlibabaPricePromotionItemDeleteResult 
type AlibabaPricePromotionItemDeleteResult struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // data
    DataList   []string `json:"data_list,omitempty"`

    // code
    ResultCode   int64 `json:"result_code,omitempty"`

    // totalRecord
    TotalRecord   int64 `json:"total_record,omitempty"`

    // msg
    Message   string `json:"message,omitempty"`

}
