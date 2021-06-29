package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版查询 API请求
alibaba.wdk.coupon.template.query

优惠券模版查询
*/
type AlibabaWdkCouponTemplateQueryRequest struct {
    model.Params
    // 系统自动生成
    _paramCouponTemplateQueryRequest   *CouponTemplateQueryRequest
}

// 初始化AlibabaWdkCouponTemplateQueryRequest对象
func NewAlibabaWdkCouponTemplateQueryRequest() *AlibabaWdkCouponTemplateQueryRequest{
    return &AlibabaWdkCouponTemplateQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.template.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCouponTemplateQueryRequest Setter
// 系统自动生成
func (r *AlibabaWdkCouponTemplateQueryRequest) SetParamCouponTemplateQueryRequest(_paramCouponTemplateQueryRequest *CouponTemplateQueryRequest) error {
    r._paramCouponTemplateQueryRequest = _paramCouponTemplateQueryRequest
    r.Set("param_coupon_template_query_request", _paramCouponTemplateQueryRequest)
    return nil
}

// ParamCouponTemplateQueryRequest Getter
func (r AlibabaWdkCouponTemplateQueryRequest) GetParamCouponTemplateQueryRequest() *CouponTemplateQueryRequest {
    return r._paramCouponTemplateQueryRequest
}
