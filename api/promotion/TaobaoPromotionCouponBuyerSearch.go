package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询买家在相关app领取的优惠券信息 
taobao.promotion.coupon.buyer.search

查询买家在相关app领取的优惠券信息
*/
func TaobaoPromotionCouponBuyerSearch(clt *core.SDKClient, req *promotion.TaobaoPromotionCouponBuyerSearchAPIRequest, session string) (*promotion.TaobaoPromotionCouponBuyerSearchAPIResponse, error) {
    var resp promotion.TaobaoPromotionCouponBuyerSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
