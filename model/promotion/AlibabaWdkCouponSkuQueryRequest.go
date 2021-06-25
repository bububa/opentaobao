package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品查询 APIRequest
alibaba.wdk.coupon.sku.query

优惠券商品查询
*/
type AlibabaWdkCouponSkuQueryRequest struct {
    model.Params

    // 请求
    paramCouponTemplateItemQueryRequest   *CouponTemplateItemQueryRequest 

}

func NewAlibabaWdkCouponSkuQueryRequest() *AlibabaWdkCouponSkuQueryRequest{
    return &AlibabaWdkCouponSkuQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkCouponSkuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.sku.query"
}

func (r AlibabaWdkCouponSkuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkCouponSkuQueryRequest) SetParamCouponTemplateItemQueryRequest(paramCouponTemplateItemQueryRequest *CouponTemplateItemQueryRequest) error {
    r.paramCouponTemplateItemQueryRequest = paramCouponTemplateItemQueryRequest
    r.Set("param_coupon_template_item_query_request", paramCouponTemplateItemQueryRequest)
    return nil
}

func (r AlibabaWdkCouponSkuQueryRequest) GetParamCouponTemplateItemQueryRequest() *CouponTemplateItemQueryRequest {
    return r.paramCouponTemplateItemQueryRequest
}

