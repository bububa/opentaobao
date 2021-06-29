package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品查询 API请求
alibaba.wdk.coupon.sku.query

优惠券商品查询
*/
type AlibabaWdkCouponSkuQueryRequest struct {
    model.Params
    // 请求
    _paramCouponTemplateItemQueryRequest   *CouponTemplateItemQueryRequest
}

// 初始化AlibabaWdkCouponSkuQueryRequest对象
func NewAlibabaWdkCouponSkuQueryRequest() *AlibabaWdkCouponSkuQueryRequest{
    return &AlibabaWdkCouponSkuQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSkuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.sku.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSkuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCouponTemplateItemQueryRequest Setter
// 请求
func (r *AlibabaWdkCouponSkuQueryRequest) SetParamCouponTemplateItemQueryRequest(_paramCouponTemplateItemQueryRequest *CouponTemplateItemQueryRequest) error {
    r._paramCouponTemplateItemQueryRequest = _paramCouponTemplateItemQueryRequest
    r.Set("param_coupon_template_item_query_request", _paramCouponTemplateItemQueryRequest)
    return nil
}

// ParamCouponTemplateItemQueryRequest Getter
func (r AlibabaWdkCouponSkuQueryRequest) GetParamCouponTemplateItemQueryRequest() *CouponTemplateItemQueryRequest {
    return r._paramCouponTemplateItemQueryRequest
}
