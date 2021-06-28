package jst

// TaobaoRdsDbCreateaccountResultSet 
/* model for simplify = false
type TaobaoRdsDbCreateaccountResultSet struct {

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

// TaobaoRdsDbCreateaccountResultSet 
type TaobaoRdsDbCreateaccountResultSet struct {

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
