package promotion

// TmallPromotionCouponUserResult 
/* model for simplify = false
type TmallPromotionCouponUserResult struct {

    // resultCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // data
    
    Data  *struct {
        UserInfoDo  *UserInfoDo `json:"user_info_do,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

// TmallPromotionCouponUserResult 
type TmallPromotionCouponUserResult struct {

    // resultCode
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // data
    Data   *UserInfoDo `json:"data,omitempty"`

}
