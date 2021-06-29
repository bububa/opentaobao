package xhotelitem

// TaobaoXhotelServicetimeGetResultSet 
type TaobaoXhotelServicetimeGetResultSet struct {
    // results
    Results   []string `json:"results,omitempty" xml:"results>string,omitempty"`
    // totalResults
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
    // hasNext
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
    // exception
    Exception   string `json:"exception,omitempty" xml:"exception,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // warnMessage
    WarnMessage   string `json:"warn_message,omitempty" xml:"warn_message,omitempty"`
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // firstResult
    FirstResults   []ServiceTimeDataDO `json:"first_results,omitempty" xml:"first_results>service_time_data_do,omitempty"`
}
