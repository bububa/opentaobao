package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询绑定卖家优惠券相关信息(手淘专用) APIResponse
taobao.mobile.promotion.coupon.seller.search

查询绑定卖家相关优惠券信息 如isv 百川 等外部业务方
*/
type TaobaoMobilePromotionCouponSellerSearchAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMobilePromotionCouponSellerSearchResponse `json:"mobile_promotion_coupon_seller_search_response,omitempty"` 
    TaobaoMobilePromotionCouponSellerSearchResponse
}

/* model for simplify = false
type TaobaoMobilePromotionCouponSellerSearchResponse struct {

    // 优惠券查询结果
    
    CouponSearchResult  *struct {
        CouponSearchResult  *CouponSearchResult `json:"coupon_search_result,omitempty"`
    } `json:"coupon_search_result,omitempty"`
    

}
*/

type TaobaoMobilePromotionCouponSellerSearchResponse struct {

    // 优惠券查询结果
    CouponSearchResult   *CouponSearchResult `json:"coupon_search_result,omitempty"`

}
