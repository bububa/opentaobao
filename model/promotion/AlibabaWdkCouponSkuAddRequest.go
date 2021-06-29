package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品增加 API请求
alibaba.wdk.coupon.sku.add

优惠券商品增加
*/
type AlibabaWdkCouponSkuAddRequest struct {
    model.Params
    // 请求
    paramCouponTemplateItemRequest   *CouponTemplateItemRequest
}

// 初始化AlibabaWdkCouponSkuAddRequest对象
func NewAlibabaWdkCouponSkuAddRequest() *AlibabaWdkCouponSkuAddRequest{
    return &AlibabaWdkCouponSkuAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSkuAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCouponTemplateItemRequest Setter
// 请求
func (r *AlibabaWdkCouponSkuAddRequest) SetParamCouponTemplateItemRequest(paramCouponTemplateItemRequest *CouponTemplateItemRequest) error {
    r.paramCouponTemplateItemRequest = paramCouponTemplateItemRequest
    r.Set("param_coupon_template_item_request", paramCouponTemplateItemRequest)
    return nil
}

// ParamCouponTemplateItemRequest Getter
func (r AlibabaWdkCouponSkuAddRequest) GetParamCouponTemplateItemRequest() *CouponTemplateItemRequest {
    return r.paramCouponTemplateItemRequest
}
