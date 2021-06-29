package promotion

// TmallPromotionCouponQueryResult 
type TmallPromotionCouponQueryResult struct {
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // resultCode
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // data
    DataList   []TmallPromotionCouponQueryData `json:"data_list,omitempty" xml:"data_list>tmall_promotion_coupon_query_data,omitempty"`
}
