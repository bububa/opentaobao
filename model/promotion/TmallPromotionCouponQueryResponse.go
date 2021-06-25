package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询可用优惠券列表 APIResponse
tmall.promotion.coupon.query

查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick
*/
type TmallPromotionCouponQueryAPIResponse struct {
    model.CommonResponse
    Response *TmallPromotionCouponQueryResponse `json:"tmall_promotion_coupon_query_response,omitempty"`
}

type TmallPromotionCouponQueryResponse struct {

    // result
    Result   *TmallPromotionCouponQueryResult `json:"result,omitempty"`

}
