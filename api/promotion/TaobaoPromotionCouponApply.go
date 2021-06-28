package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
优惠券领取 
taobao.promotion.coupon.apply

优惠券领取
*/
func TaobaoPromotionCouponApply(clt *core.SDKClient, req *promotion.TaobaoPromotionCouponApplyRequest, session string) (*promotion.TaobaoPromotionCouponApplyAPIResponse, error) {
    var resp promotion.TaobaoPromotionCouponApplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
