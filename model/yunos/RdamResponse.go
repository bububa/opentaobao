package yunos

// RdamResponse 
/* model for simplify = false
type RdamResponse struct {

    // message
    
    Message   string `json:"message,omitempty"`
    

    // more
    
    More   string `json:"more,omitempty"`
    

    // code
    
    Code   int64 `json:"code,omitempty"`
    

    // dataList
    
    DataList  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"data_list,omitempty"`
    

    // traceId
    
    TraceId   string `json:"trace_id,omitempty"`
    

}
*/

// RdamResponse 
type RdamResponse struct {

    // message
    Message   string `json:"message,omitempty"`

    // more
    More   string `json:"more,omitempty"`

    // code
    Code   int64 `json:"code,omitempty"`

    // dataList
    DataList   []string `json:"data_list,omitempty"`

    // traceId
    TraceId   string `json:"trace_id,omitempty"`

}
