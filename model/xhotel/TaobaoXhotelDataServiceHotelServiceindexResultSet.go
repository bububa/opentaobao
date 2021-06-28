package xhotel

// TaobaoXhotelDataServiceHotelServiceindexResultSet 
/* model for simplify = false
type TaobaoXhotelDataServiceHotelServiceindexResultSet struct {

    // firstResult
    
    FirstResult  *struct {
        TopAdsHtlDataQueryResult  *TopAdsHtlDataQueryResult `json:"top_ads_htl_data_query_result,omitempty"`
    } `json:"first_result,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoXhotelDataServiceHotelServiceindexResultSet 
type TaobaoXhotelDataServiceHotelServiceindexResultSet struct {

    // firstResult
    FirstResult   *TopAdsHtlDataQueryResult `json:"first_result,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
