package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版终止 API请求
alibaba.wdk.coupon.template.terminate

优惠券模版终止
*/
type AlibabaWdkCouponTemplateTerminateRequest struct {
    model.Params
    // 参数
    _paramCouponTemplateTerminateRequest   *CouponTemplateTerminateRequest
}

// 初始化AlibabaWdkCouponTemplateTerminateRequest对象
func NewAlibabaWdkCouponTemplateTerminateRequest() *AlibabaWdkCouponTemplateTerminateRequest{
    return &AlibabaWdkCouponTemplateTerminateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateTerminateRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.template.terminate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateTerminateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCouponTemplateTerminateRequest Setter
// 参数
func (r *AlibabaWdkCouponTemplateTerminateRequest) SetParamCouponTemplateTerminateRequest(_paramCouponTemplateTerminateRequest *CouponTemplateTerminateRequest) error {
    r._paramCouponTemplateTerminateRequest = _paramCouponTemplateTerminateRequest
    r.Set("param_coupon_template_terminate_request", _paramCouponTemplateTerminateRequest)
    return nil
}

// ParamCouponTemplateTerminateRequest Getter
func (r AlibabaWdkCouponTemplateTerminateRequest) GetParamCouponTemplateTerminateRequest() *CouponTemplateTerminateRequest {
    return r._paramCouponTemplateTerminateRequest
}
