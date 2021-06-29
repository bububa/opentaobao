package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版创建 API请求
alibaba.wdk.coupon.template.create

开放给外部商家创建优惠券模版
*/
type AlibabaWdkCouponTemplateCreateRequest struct {
    model.Params
    // 请求
    paramCouponTemplateOperateRequest   *CouponTemplateOperateRequest
}

// 初始化AlibabaWdkCouponTemplateCreateRequest对象
func NewAlibabaWdkCouponTemplateCreateRequest() *AlibabaWdkCouponTemplateCreateRequest{
    return &AlibabaWdkCouponTemplateCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.template.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCouponTemplateOperateRequest Setter
// 请求
func (r *AlibabaWdkCouponTemplateCreateRequest) SetParamCouponTemplateOperateRequest(paramCouponTemplateOperateRequest *CouponTemplateOperateRequest) error {
    r.paramCouponTemplateOperateRequest = paramCouponTemplateOperateRequest
    r.Set("param_coupon_template_operate_request", paramCouponTemplateOperateRequest)
    return nil
}

// ParamCouponTemplateOperateRequest Getter
func (r AlibabaWdkCouponTemplateCreateRequest) GetParamCouponTemplateOperateRequest() *CouponTemplateOperateRequest {
    return r.paramCouponTemplateOperateRequest
}
