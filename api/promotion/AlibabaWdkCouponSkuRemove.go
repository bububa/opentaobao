package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
优惠券商品删除 
alibaba.wdk.coupon.sku.remove

优惠券商品删除
*/
func AlibabaWdkCouponSkuRemove(clt *core.SDKClient, req *promotion.AlibabaWdkCouponSkuRemoveRequest, session string) (*promotion.AlibabaWdkCouponSkuRemoveResponse, error) {
    var resp promotion.AlibabaWdkCouponSkuRemoveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
