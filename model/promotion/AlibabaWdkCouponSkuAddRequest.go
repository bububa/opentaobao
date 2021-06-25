package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品增加 APIRequest
alibaba.wdk.coupon.sku.add

优惠券商品增加
*/
type AlibabaWdkCouponSkuAddRequest struct {
    model.Params

    // 请求
    paramCouponTemplateItemRequest   *CouponTemplateItemRequest 

}

func NewAlibabaWdkCouponSkuAddRequest() *AlibabaWdkCouponSkuAddRequest{
    return &AlibabaWdkCouponSkuAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkCouponSkuAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.sku.add"
}

func (r AlibabaWdkCouponSkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkCouponSkuAddRequest) SetParamCouponTemplateItemRequest(paramCouponTemplateItemRequest *CouponTemplateItemRequest) error {
    r.paramCouponTemplateItemRequest = paramCouponTemplateItemRequest
    r.Set("param_coupon_template_item_request", paramCouponTemplateItemRequest)
    return nil
}

func (r AlibabaWdkCouponSkuAddRequest) GetParamCouponTemplateItemRequest() *CouponTemplateItemRequest {
    return r.paramCouponTemplateItemRequest
}

