package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
创建店铺优惠券接口 
taobao.promotion.coupon.add

创建店铺优惠券。有效期内的店铺优惠券总数量不超过50张
*/
func TaobaoPromotionCouponAdd(clt *core.SDKClient, req *promotion.TaobaoPromotionCouponAddAPIRequest, session string) (*promotion.TaobaoPromotionCouponAddAPIResponse, error) {
    var resp promotion.TaobaoPromotionCouponAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
