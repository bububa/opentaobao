package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询绑定卖家优惠券相关信息 
taobao.promotion.coupon.seller.search

查询绑定卖家相关优惠券信息  如isv  百川 等外部业务方
*/
func TaobaoPromotionCouponSellerSearch(clt *core.SDKClient, req *promotion.TaobaoPromotionCouponSellerSearchAPIRequest, session string) (*promotion.TaobaoPromotionCouponSellerSearchAPIResponse, error) {
    var resp promotion.TaobaoPromotionCouponSellerSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
