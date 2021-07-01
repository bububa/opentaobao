package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
优惠券领取(手淘专用) 
taobao.mobile.promotion.coupon.apply

优惠券领取
*/
func TaobaoMobilePromotionCouponApply(clt *core.SDKClient, req *promotion.TaobaoMobilePromotionCouponApplyAPIRequest, session string) (*promotion.TaobaoMobilePromotionCouponApplyAPIResponse, error) {
    var resp promotion.TaobaoMobilePromotionCouponApplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
