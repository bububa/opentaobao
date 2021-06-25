package mos

// MultiResult 
type MultiResult struct {

    // total
    ResultTotal   int64 `json:"result_total,omitempty"`

    // errMessage
    ResultMessage   string `json:"result_message,omitempty"`

    // data
    ResultDatas   []Json `json:"result_datas,omitempty"`

    // errCode
    ResultCode   string `json:"result_code,omitempty"`

    // success
    ResultSuccess   bool `json:"result_success,omitempty"`

}
