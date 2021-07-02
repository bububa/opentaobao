package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponSkuRemoveAPIRequest 优惠券商品删除 API请求
// alibaba.wdk.coupon.sku.remove
//
// 优惠券商品删除
type AlibabaWdkCouponSkuRemoveAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemRequest *CouponTemplateItemRequest
}

// NewAlibabaWdkCouponSkuRemoveRequest 初始化AlibabaWdkCouponSkuRemoveAPIRequest对象
func NewAlibabaWdkCouponSkuRemoveRequest() *AlibabaWdkCouponSkuRemoveAPIRequest {
	return &AlibabaWdkCouponSkuRemoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSkuRemoveAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.sku.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSkuRemoveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamCouponTemplateItemRequest is ParamCouponTemplateItemRequest Setter
// 请求
func (r *AlibabaWdkCouponSkuRemoveAPIRequest) SetParamCouponTemplateItemRequest(_paramCouponTemplateItemRequest *CouponTemplateItemRequest) error {
	r._paramCouponTemplateItemRequest = _paramCouponTemplateItemRequest
	r.Set("param_coupon_template_item_request", _paramCouponTemplateItemRequest)
	return nil
}

// GetParamCouponTemplateItemRequest ParamCouponTemplateItemRequest Getter
func (r AlibabaWdkCouponSkuRemoveAPIRequest) GetParamCouponTemplateItemRequest() *CouponTemplateItemRequest {
	return r._paramCouponTemplateItemRequest
}
