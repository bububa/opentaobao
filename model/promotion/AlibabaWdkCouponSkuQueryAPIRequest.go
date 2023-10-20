package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponSkuQueryAPIRequest 优惠券商品查询 API请求
// alibaba.wdk.coupon.sku.query
//
// 优惠券商品查询
type AlibabaWdkCouponSkuQueryAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemQueryRequest *CouponTemplateItemQueryRequest
}

// NewAlibabaWdkCouponSkuQueryRequest 初始化AlibabaWdkCouponSkuQueryAPIRequest对象
func NewAlibabaWdkCouponSkuQueryRequest() *AlibabaWdkCouponSkuQueryAPIRequest {
	return &AlibabaWdkCouponSkuQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkCouponSkuQueryAPIRequest) Reset() {
	r._paramCouponTemplateItemQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSkuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSkuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkCouponSkuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateItemQueryRequest is ParamCouponTemplateItemQueryRequest Setter
// 请求
func (r *AlibabaWdkCouponSkuQueryAPIRequest) SetParamCouponTemplateItemQueryRequest(_paramCouponTemplateItemQueryRequest *CouponTemplateItemQueryRequest) error {
	r._paramCouponTemplateItemQueryRequest = _paramCouponTemplateItemQueryRequest
	r.Set("param_coupon_template_item_query_request", _paramCouponTemplateItemQueryRequest)
	return nil
}

// GetParamCouponTemplateItemQueryRequest ParamCouponTemplateItemQueryRequest Getter
func (r AlibabaWdkCouponSkuQueryAPIRequest) GetParamCouponTemplateItemQueryRequest() *CouponTemplateItemQueryRequest {
	return r._paramCouponTemplateItemQueryRequest
}

var poolAlibabaWdkCouponSkuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkCouponSkuQueryRequest()
	},
}

// GetAlibabaWdkCouponSkuQueryRequest 从 sync.Pool 获取 AlibabaWdkCouponSkuQueryAPIRequest
func GetAlibabaWdkCouponSkuQueryAPIRequest() *AlibabaWdkCouponSkuQueryAPIRequest {
	return poolAlibabaWdkCouponSkuQueryAPIRequest.Get().(*AlibabaWdkCouponSkuQueryAPIRequest)
}

// ReleaseAlibabaWdkCouponSkuQueryAPIRequest 将 AlibabaWdkCouponSkuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkCouponSkuQueryAPIRequest(v *AlibabaWdkCouponSkuQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkCouponSkuQueryAPIRequest.Put(v)
}
