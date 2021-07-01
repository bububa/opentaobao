package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询可用优惠券列表 
tmall.promotion.coupon.query

查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick
*/
func TmallPromotionCouponQuery(clt *core.SDKClient, req *promotion.TmallPromotionCouponQueryAPIRequest, session string) (*promotion.TmallPromotionCouponQueryAPIResponse, error) {
    var resp promotion.TmallPromotionCouponQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
