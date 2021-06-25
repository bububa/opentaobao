package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券模版查询 APIRequest
alibaba.wdk.coupon.template.query

优惠券模版查询
*/
type AlibabaWdkCouponTemplateQueryRequest struct {
    model.Params

    // 系统自动生成
    paramCouponTemplateQueryRequest   *CouponTemplateQueryRequest 

}

func NewAlibabaWdkCouponTemplateQueryRequest() *AlibabaWdkCouponTemplateQueryRequest{
    return &AlibabaWdkCouponTemplateQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkCouponTemplateQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.template.query"
}

func (r AlibabaWdkCouponTemplateQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkCouponTemplateQueryRequest) SetParamCouponTemplateQueryRequest(paramCouponTemplateQueryRequest *CouponTemplateQueryRequest) error {
    r.paramCouponTemplateQueryRequest = paramCouponTemplateQueryRequest
    r.Set("param_coupon_template_query_request", paramCouponTemplateQueryRequest)
    return nil
}

func (r AlibabaWdkCouponTemplateQueryRequest) GetParamCouponTemplateQueryRequest() *CouponTemplateQueryRequest {
    return r.paramCouponTemplateQueryRequest
}

