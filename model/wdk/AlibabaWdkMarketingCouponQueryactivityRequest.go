package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询优惠券活动 APIRequest
alibaba.wdk.marketing.coupon.queryactivity

查询优惠券活动
*/
type AlibabaWdkMarketingCouponQueryactivityRequest struct {
    model.Params

    // 查询优惠券活动入参
    param0   *CommonActivityParam 

}

func NewAlibabaWdkMarketingCouponQueryactivityRequest() *AlibabaWdkMarketingCouponQueryactivityRequest{
    return &AlibabaWdkMarketingCouponQueryactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingCouponQueryactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.queryactivity"
}

func (r AlibabaWdkMarketingCouponQueryactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingCouponQueryactivityRequest) SetParam0(param0 *CommonActivityParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingCouponQueryactivityRequest) GetParam0() *CommonActivityParam {
    return r.param0
}

