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
type AlibabaWdkCouponSkuRemoveAPIRequest struct {
    model.Params
    // 请求
    _paramCouponTemplateItemRequest   *CouponTemplateItemRequest
}

// 初始化AlibabaWdkCouponSkuRemoveAPIRequest对象
func NewAlibabaWdkCouponSkuRemoveRequest() *AlibabaWdkCouponSkuRemoveAPIRequest{
    return &AlibabaWdkCouponSkuRemoveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSkuRemoveAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.sku.remove"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSkuRemoveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCouponTemplateItemRequest Setter
// 请求
func (r *AlibabaWdkCouponSkuRemoveAPIRequest) SetParamCouponTemplateItemRequest(_paramCouponTemplateItemRequest *CouponTemplateItemRequest) error {
    r._paramCouponTemplateItemRequest = _paramCouponTemplateItemRequest
    r.Set("param_coupon_template_item_request", _paramCouponTemplateItemRequest)
    return nil
}

// ParamCouponTemplateItemRequest Getter
func (r AlibabaWdkCouponSkuRemoveAPIRequest) GetParamCouponTemplateItemRequest() *CouponTemplateItemRequest {
    return r._paramCouponTemplateItemRequest
}
