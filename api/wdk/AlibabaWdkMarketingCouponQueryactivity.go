package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查询优惠券活动 
alibaba.wdk.marketing.coupon.queryactivity

查询优惠券活动
*/
func AlibabaWdkMarketingCouponQueryactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingCouponQueryactivityRequest, session string) (*wdk.AlibabaWdkMarketingCouponQueryactivityResponse, error) {
    var resp wdk.AlibabaWdkMarketingCouponQueryactivityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
