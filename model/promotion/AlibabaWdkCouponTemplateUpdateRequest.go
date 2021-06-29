package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版修改 API请求
alibaba.wdk.coupon.template.update

优惠券模版修改
*/
type AlibabaWdkCouponTemplateUpdateRequest struct {
    model.Params
    // 请求
    _paramCouponTemplateOperateRequest   *CouponTemplateOperateRequest
}

// 初始化AlibabaWdkCouponTemplateUpdateRequest对象
func NewAlibabaWdkCouponTemplateUpdateRequest() *AlibabaWdkCouponTemplateUpdateRequest{
    return &AlibabaWdkCouponTemplateUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.template.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCouponTemplateOperateRequest Setter
// 请求
func (r *AlibabaWdkCouponTemplateUpdateRequest) SetParamCouponTemplateOperateRequest(_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest) error {
    r._paramCouponTemplateOperateRequest = _paramCouponTemplateOperateRequest
    r.Set("param_coupon_template_operate_request", _paramCouponTemplateOperateRequest)
    return nil
}

// ParamCouponTemplateOperateRequest Getter
func (r AlibabaWdkCouponTemplateUpdateRequest) GetParamCouponTemplateOperateRequest() *CouponTemplateOperateRequest {
    return r._paramCouponTemplateOperateRequest
}
