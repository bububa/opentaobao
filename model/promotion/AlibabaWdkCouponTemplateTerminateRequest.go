package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版终止 APIRequest
alibaba.wdk.coupon.template.terminate

优惠券模版终止
*/
type AlibabaWdkCouponTemplateTerminateRequest struct {
    model.Params

    // 参数
    paramCouponTemplateTerminateRequest   *CouponTemplateTerminateRequest 

}

func NewAlibabaWdkCouponTemplateTerminateRequest() *AlibabaWdkCouponTemplateTerminateRequest{
    return &AlibabaWdkCouponTemplateTerminateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkCouponTemplateTerminateRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.template.terminate"
}

func (r AlibabaWdkCouponTemplateTerminateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkCouponTemplateTerminateRequest) SetParamCouponTemplateTerminateRequest(paramCouponTemplateTerminateRequest *CouponTemplateTerminateRequest) error {
    r.paramCouponTemplateTerminateRequest = paramCouponTemplateTerminateRequest
    r.Set("param_coupon_template_terminate_request", paramCouponTemplateTerminateRequest)
    return nil
}

func (r AlibabaWdkCouponTemplateTerminateRequest) GetParamCouponTemplateTerminateRequest() *CouponTemplateTerminateRequest {
    return r.paramCouponTemplateTerminateRequest
}

