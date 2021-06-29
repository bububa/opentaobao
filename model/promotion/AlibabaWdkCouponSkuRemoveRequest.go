package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品删除 API请求
alibaba.wdk.coupon.sku.remove

优惠券商品删除
*/
type AlibabaWdkCouponSkuRemoveRequest struct {
    model.Params
    // 请求
    paramCouponTemplateItemRequest   *CouponTemplateItemRequest
}

// 初始化AlibabaWdkCouponSkuRemoveRequest对象
func NewAlibabaWdkCouponSkuRemoveRequest() *AlibabaWdkCouponSkuRemoveRequest{
    return &AlibabaWdkCouponSkuRemoveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSkuRemoveRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.sku.remove"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSkuRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCouponTemplateItemRequest Setter
// 请求
func (r *AlibabaWdkCouponSkuRemoveRequest) SetParamCouponTemplateItemRequest(paramCouponTemplateItemRequest *CouponTemplateItemRequest) error {
    r.paramCouponTemplateItemRequest = paramCouponTemplateItemRequest
    r.Set("param_coupon_template_item_request", paramCouponTemplateItemRequest)
    return nil
}

// ParamCouponTemplateItemRequest Getter
func (r AlibabaWdkCouponSkuRemoveRequest) GetParamCouponTemplateItemRequest() *CouponTemplateItemRequest {
    return r.paramCouponTemplateItemRequest
}
