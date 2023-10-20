package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponSkuAddAPIRequest 优惠券商品增加 API请求
// alibaba.wdk.coupon.sku.add
//
// 优惠券商品增加
type AlibabaWdkCouponSkuAddAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemRequest *CouponTemplateItemRequest
}

// NewAlibabaWdkCouponSkuAddRequest 初始化AlibabaWdkCouponSkuAddAPIRequest对象
func NewAlibabaWdkCouponSkuAddRequest() *AlibabaWdkCouponSkuAddAPIRequest {
	return &AlibabaWdkCouponSkuAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkCouponSkuAddAPIRequest) Reset() {
	r._paramCouponTemplateItemRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSkuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSkuAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkCouponSkuAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateItemRequest is ParamCouponTemplateItemRequest Setter
// 请求
func (r *AlibabaWdkCouponSkuAddAPIRequest) SetParamCouponTemplateItemRequest(_paramCouponTemplateItemRequest *CouponTemplateItemRequest) error {
	r._paramCouponTemplateItemRequest = _paramCouponTemplateItemRequest
	r.Set("param_coupon_template_item_request", _paramCouponTemplateItemRequest)
	return nil
}

// GetParamCouponTemplateItemRequest ParamCouponTemplateItemRequest Getter
func (r AlibabaWdkCouponSkuAddAPIRequest) GetParamCouponTemplateItemRequest() *CouponTemplateItemRequest {
	return r._paramCouponTemplateItemRequest
}

var poolAlibabaWdkCouponSkuAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkCouponSkuAddRequest()
	},
}

// GetAlibabaWdkCouponSkuAddRequest 从 sync.Pool 获取 AlibabaWdkCouponSkuAddAPIRequest
func GetAlibabaWdkCouponSkuAddAPIRequest() *AlibabaWdkCouponSkuAddAPIRequest {
	return poolAlibabaWdkCouponSkuAddAPIRequest.Get().(*AlibabaWdkCouponSkuAddAPIRequest)
}

// ReleaseAlibabaWdkCouponSkuAddAPIRequest 将 AlibabaWdkCouponSkuAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkCouponSkuAddAPIRequest(v *AlibabaWdkCouponSkuAddAPIRequest) {
	v.Reset()
	poolAlibabaWdkCouponSkuAddAPIRequest.Put(v)
}
