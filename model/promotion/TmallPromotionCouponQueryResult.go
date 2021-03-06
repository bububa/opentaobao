package promotion

// TmallPromotionCouponQueryResult 结构体
type TmallPromotionCouponQueryResult struct {
	// data
	DataList []TmallPromotionCouponQueryData `json:"data_list,omitempty" xml:"data_list>tmall_promotion_coupon_query_data,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// resultCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
