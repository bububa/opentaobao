package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发放匿名码 APIRequest
alibaba.wdk.marketing.coupon.sendma

根据优惠券活动id打印单个匿名码
*/
type AlibabaWdkMarketingCouponSendmaRequest struct {
    model.Params

    // 发放匿名码入参
    param0   *CommonActivityParam 

}

func NewAlibabaWdkMarketingCouponSendmaRequest() *AlibabaWdkMarketingCouponSendmaRequest{
    return &AlibabaWdkMarketingCouponSendmaRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingCouponSendmaRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.sendma"
}

func (r AlibabaWdkMarketingCouponSendmaRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingCouponSendmaRequest) SetParam0(param0 *CommonActivityParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkMarketingCouponSendmaRequest) GetParam0() *CommonActivityParam {
    return r.param0
}

