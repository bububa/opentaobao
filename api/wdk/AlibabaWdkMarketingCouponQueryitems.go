package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查询优惠券活动下的商品 
alibaba.wdk.marketing.coupon.queryitems

查询优惠券活动下面的商品
*/
func AlibabaWdkMarketingCouponQueryitems(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingCouponQueryitemsRequest, session string) (*wdk.AlibabaWdkMarketingCouponQueryitemsAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingCouponQueryitemsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
