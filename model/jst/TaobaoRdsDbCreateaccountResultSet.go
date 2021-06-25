package jst

// TaobaoRdsDbCreateaccountResultSet 
type TaobaoRdsDbCreateaccountResultSet struct {

    // results
    Results   []Json `json:"results,omitempty"`

    // totalResults
    TotalResults   int64 `json:"total_results,omitempty"`

    // exception
    Exception   string `json:"exception,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

}
