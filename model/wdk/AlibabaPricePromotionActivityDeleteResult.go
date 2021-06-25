package wdk

// AlibabaPricePromotionActivityDeleteResult 
type AlibabaPricePromotionActivityDeleteResult struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // data
    DataList   []String `json:"data_list,omitempty"`

    // code
    ResultCode   int64 `json:"result_code,omitempty"`

    // totalRecord
    TotalRecord   int64 `json:"total_record,omitempty"`

    // msg
    Message   string `json:"message,omitempty"`

}
