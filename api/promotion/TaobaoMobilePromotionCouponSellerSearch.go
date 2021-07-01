package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询绑定卖家优惠券相关信息(手淘专用) 
taobao.mobile.promotion.coupon.seller.search

查询绑定卖家相关优惠券信息 如isv 百川 等外部业务方
*/
func TaobaoMobilePromotionCouponSellerSearch(clt *core.SDKClient, req *promotion.TaobaoMobilePromotionCouponSellerSearchAPIRequest, session string) (*promotion.TaobaoMobilePromotionCouponSellerSearchAPIResponse, error) {
    var resp promotion.TaobaoMobilePromotionCouponSellerSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
