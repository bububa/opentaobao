package jst

// TaobaoRdsDbGetdbResultSet 
/* model for simplify = false
type TaobaoRdsDbGetdbResultSet struct {

    // results
    
    Results  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"results,omitempty"`
    

    // totalResults
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // exception
    
    Exception   string `json:"exception,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// TaobaoRdsDbGetdbResultSet 
type TaobaoRdsDbGetdbResultSet struct {

    // results
    Results   []string `json:"results,omitempty"`

    // totalResults
    TotalResults   int64 `json:"total_results,omitempty"`

    // exception
    Exception   string `json:"exception,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

}
