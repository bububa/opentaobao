package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponSkuAddAPIRequest
优惠券商品增加 API请求
alibaba.wdk.coupon.sku.add

优惠券商品增加 */
type AlibabaWdkCouponSkuAddAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemRequest *CouponTemplateItemRequest
}

// NewAlibabaWdkCouponSkuAddRequest 初始化AlibabaWdkCouponSkuAddAPIRequest对象
func NewAlibabaWdkCouponSkuAddRequest() *AlibabaWdkCouponSkuAddAPIRequest {
	return &AlibabaWdkCouponSkuAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSkuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSkuAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamCouponTemplateItemRequest Setter
// 请求
func (r *AlibabaWdkCouponSkuAddAPIRequest) SetParamCouponTemplateItemRequest(_paramCouponTemplateItemRequest *CouponTemplateItemRequest) error {
	r._paramCouponTemplateItemRequest = _paramCouponTemplateItemRequest
	r.Set("param_coupon_template_item_request", _paramCouponTemplateItemRequest)
	return nil
}

// Get ParamCouponTemplateItemRequest Getter
func (r AlibabaWdkCouponSkuAddAPIRequest) GetParamCouponTemplateItemRequest() *CouponTemplateItemRequest {
	return r._paramCouponTemplateItemRequest
}
