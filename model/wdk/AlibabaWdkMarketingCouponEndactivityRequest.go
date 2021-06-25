package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
结束优惠券活动 APIRequest
alibaba.wdk.marketing.coupon.endactivity

结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用
*/
type AlibabaWdkMarketingCouponEndactivityRequest struct {
    model.Params

    // 需要删除的活动的信息
    param   *CommonActivityParam 

}

func NewAlibabaWdkMarketingCouponEndactivityRequest() *AlibabaWdkMarketingCouponEndactivityRequest{
    return &AlibabaWdkMarketingCouponEndactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingCouponEndactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.endactivity"
}

func (r AlibabaWdkMarketingCouponEndactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingCouponEndactivityRequest) SetParam(param *CommonActivityParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingCouponEndactivityRequest) GetParam() *CommonActivityParam {
    return r.param
}

