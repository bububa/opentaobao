package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
券核销接口 APIResponse
tmall.promotion.coupon.use

核销用户的一张优惠券，返回核销结果
*/
type TmallPromotionCouponUseAPIResponse struct {
    model.CommonResponse
    // Response *TmallPromotionCouponUseResponse `json:"tmall_promotion_coupon_use_response,omitempty"` 
    TmallPromotionCouponUseResponse
}

/* model for simplify = false
type TmallPromotionCouponUseResponse struct {

    // data
    
    Data  *struct {
        UseResultDo  *UseResultDo `json:"use_result_do,omitempty"`
    } `json:"data,omitempty"`
    

    // resultCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

type TmallPromotionCouponUseResponse struct {

    // data
    Data   *UseResultDo `json:"data,omitempty"`

    // resultCode
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

}
