package promotion

// TmallPromotionCouponUserResult 
type TmallPromotionCouponUserResult struct {

    // resultCode
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // data
    Data   *UserInfoDo `json:"data,omitempty"`

}
