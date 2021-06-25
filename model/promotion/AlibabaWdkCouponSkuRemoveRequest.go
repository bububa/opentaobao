package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品删除 APIRequest
alibaba.wdk.coupon.sku.remove

优惠券商品删除
*/
type AlibabaWdkCouponSkuRemoveRequest struct {
    model.Params

    // 请求
    paramCouponTemplateItemRequest   *CouponTemplateItemRequest 

}

func NewAlibabaWdkCouponSkuRemoveRequest() *AlibabaWdkCouponSkuRemoveRequest{
    return &AlibabaWdkCouponSkuRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkCouponSkuRemoveRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.sku.remove"
}

func (r AlibabaWdkCouponSkuRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkCouponSkuRemoveRequest) SetParamCouponTemplateItemRequest(paramCouponTemplateItemRequest *CouponTemplateItemRequest) error {
    r.paramCouponTemplateItemRequest = paramCouponTemplateItemRequest
    r.Set("param_coupon_template_item_request", paramCouponTemplateItemRequest)
    return nil
}

func (r AlibabaWdkCouponSkuRemoveRequest) GetParamCouponTemplateItemRequest() *CouponTemplateItemRequest {
    return r.paramCouponTemplateItemRequest
}

