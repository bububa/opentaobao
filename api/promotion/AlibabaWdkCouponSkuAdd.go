package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
优惠券商品增加 
alibaba.wdk.coupon.sku.add

优惠券商品增加
*/
func AlibabaWdkCouponSkuAdd(clt *core.SDKClient, req *promotion.AlibabaWdkCouponSkuAddRequest, session string) (*promotion.AlibabaWdkCouponSkuAddAPIResponse, error) {
    var resp promotion.AlibabaWdkCouponSkuAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
