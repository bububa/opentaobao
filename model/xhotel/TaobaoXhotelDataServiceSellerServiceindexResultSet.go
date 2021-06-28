package xhotel

// TaobaoXhotelDataServiceSellerServiceindexResultSet 
/* model for simplify = false
type TaobaoXhotelDataServiceSellerServiceindexResultSet struct {

    // firstResult
    
    FirstResult  *struct {
        TopAdsSlrQueryResult  *TopAdsSlrQueryResult `json:"top_ads_slr_query_result,omitempty"`
    } `json:"first_result,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 是否成功
    
    Success   string `json:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

}
*/

// TaobaoXhotelDataServiceSellerServiceindexResultSet 
type TaobaoXhotelDataServiceSellerServiceindexResultSet struct {

    // firstResult
    FirstResult   *TopAdsSlrQueryResult `json:"first_result,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否成功
    Success   string `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

}
