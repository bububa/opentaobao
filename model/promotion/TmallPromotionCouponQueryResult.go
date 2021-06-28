package promotion

// TmallPromotionCouponQueryResult 
/* model for simplify = false
type TmallPromotionCouponQueryResult struct {

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // resultCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // data
    
    DataList  struct {
        TmallPromotionCouponQueryData  []TmallPromotionCouponQueryData `json:"tmall_promotion_coupon_query_data,omitempty"`
    } `json:"data_list,omitempty"`
    

}
*/

// TmallPromotionCouponQueryResult 
type TmallPromotionCouponQueryResult struct {

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // resultCode
    ResultCode   string `json:"result_code,omitempty"`

    // data
    DataList   []TmallPromotionCouponQueryData `json:"data_list,omitempty"`

}
