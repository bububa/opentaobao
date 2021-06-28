package xhotel

// TaobaoXhotelDataServiceOrderDetailResultSet 
/* model for simplify = false
type TaobaoXhotelDataServiceOrderDetailResultSet struct {

    // firstResult
    
    FirstResult  *struct {
        TopAdsTripSvcQueryResult  *TopAdsTripSvcQueryResult `json:"top_ads_trip_svc_query_result,omitempty"`
    } `json:"first_result,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// TaobaoXhotelDataServiceOrderDetailResultSet 
type TaobaoXhotelDataServiceOrderDetailResultSet struct {

    // firstResult
    FirstResult   *TopAdsTripSvcQueryResult `json:"first_result,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

}
