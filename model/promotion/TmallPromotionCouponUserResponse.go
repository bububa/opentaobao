package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户信息查询接口 APIResponse
tmall.promotion.coupon.user

开发给外部合作商（例如：苏宁），通过会员付款码获得会员nick
*/
type TmallPromotionCouponUserAPIResponse struct {
    model.CommonResponse
    Response *TmallPromotionCouponUserResponse `json:"tmall_promotion_coupon_user_response,omitempty"`
}

type TmallPromotionCouponUserResponse struct {

    // result
    Result   *TmallPromotionCouponUserResult `json:"result,omitempty"`

}
