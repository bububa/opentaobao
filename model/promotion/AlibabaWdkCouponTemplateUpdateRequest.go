package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版修改 APIRequest
alibaba.wdk.coupon.template.update

优惠券模版修改
*/
type AlibabaWdkCouponTemplateUpdateRequest struct {
    model.Params

    // 请求
    paramCouponTemplateOperateRequest   *CouponTemplateOperateRequest 

}

func NewAlibabaWdkCouponTemplateUpdateRequest() *AlibabaWdkCouponTemplateUpdateRequest{
    return &AlibabaWdkCouponTemplateUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkCouponTemplateUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.template.update"
}

func (r AlibabaWdkCouponTemplateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkCouponTemplateUpdateRequest) SetParamCouponTemplateOperateRequest(paramCouponTemplateOperateRequest *CouponTemplateOperateRequest) error {
    r.paramCouponTemplateOperateRequest = paramCouponTemplateOperateRequest
    r.Set("param_coupon_template_operate_request", paramCouponTemplateOperateRequest)
    return nil
}

func (r AlibabaWdkCouponTemplateUpdateRequest) GetParamCouponTemplateOperateRequest() *CouponTemplateOperateRequest {
    return r.paramCouponTemplateOperateRequest
}

