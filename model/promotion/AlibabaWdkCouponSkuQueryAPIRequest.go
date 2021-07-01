package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponSkuQueryAPIRequest
优惠券商品查询 API请求
alibaba.wdk.coupon.sku.query

优惠券商品查询 */
type AlibabaWdkCouponSkuQueryAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemQueryRequest *CouponTemplateItemQueryRequest
}

// NewAlibabaWdkCouponSkuQueryRequest 初始化AlibabaWdkCouponSkuQueryAPIRequest对象
func NewAlibabaWdkCouponSkuQueryRequest() *AlibabaWdkCouponSkuQueryAPIRequest {
	return &AlibabaWdkCouponSkuQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSkuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSkuQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamCouponTemplateItemQueryRequest Setter
// 请求
func (r *AlibabaWdkCouponSkuQueryAPIRequest) SetParamCouponTemplateItemQueryRequest(_paramCouponTemplateItemQueryRequest *CouponTemplateItemQueryRequest) error {
	r._paramCouponTemplateItemQueryRequest = _paramCouponTemplateItemQueryRequest
	r.Set("param_coupon_template_item_query_request", _paramCouponTemplateItemQueryRequest)
	return nil
}

// Get ParamCouponTemplateItemQueryRequest Getter
func (r AlibabaWdkCouponSkuQueryAPIRequest) GetParamCouponTemplateItemQueryRequest() *CouponTemplateItemQueryRequest {
	return r._paramCouponTemplateItemQueryRequest
}
