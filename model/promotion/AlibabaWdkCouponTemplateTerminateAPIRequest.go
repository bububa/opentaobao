package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateTerminateAPIRequest 优惠券模版终止 API请求
// alibaba.wdk.coupon.template.terminate
//
// 优惠券模版终止
type AlibabaWdkCouponTemplateTerminateAPIRequest struct {
	model.Params
	// 参数
	_paramCouponTemplateTerminateRequest *CouponTemplateTerminateRequest
}

// NewAlibabaWdkCouponTemplateTerminateRequest 初始化AlibabaWdkCouponTemplateTerminateAPIRequest对象
func NewAlibabaWdkCouponTemplateTerminateRequest() *AlibabaWdkCouponTemplateTerminateAPIRequest {
	return &AlibabaWdkCouponTemplateTerminateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkCouponTemplateTerminateAPIRequest) Reset() {
	r._paramCouponTemplateTerminateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateTerminateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.terminate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateTerminateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkCouponTemplateTerminateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateTerminateRequest is ParamCouponTemplateTerminateRequest Setter
// 参数
func (r *AlibabaWdkCouponTemplateTerminateAPIRequest) SetParamCouponTemplateTerminateRequest(_paramCouponTemplateTerminateRequest *CouponTemplateTerminateRequest) error {
	r._paramCouponTemplateTerminateRequest = _paramCouponTemplateTerminateRequest
	r.Set("param_coupon_template_terminate_request", _paramCouponTemplateTerminateRequest)
	return nil
}

// GetParamCouponTemplateTerminateRequest ParamCouponTemplateTerminateRequest Getter
func (r AlibabaWdkCouponTemplateTerminateAPIRequest) GetParamCouponTemplateTerminateRequest() *CouponTemplateTerminateRequest {
	return r._paramCouponTemplateTerminateRequest
}

var poolAlibabaWdkCouponTemplateTerminateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkCouponTemplateTerminateRequest()
	},
}

// GetAlibabaWdkCouponTemplateTerminateRequest 从 sync.Pool 获取 AlibabaWdkCouponTemplateTerminateAPIRequest
func GetAlibabaWdkCouponTemplateTerminateAPIRequest() *AlibabaWdkCouponTemplateTerminateAPIRequest {
	return poolAlibabaWdkCouponTemplateTerminateAPIRequest.Get().(*AlibabaWdkCouponTemplateTerminateAPIRequest)
}

// ReleaseAlibabaWdkCouponTemplateTerminateAPIRequest 将 AlibabaWdkCouponTemplateTerminateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkCouponTemplateTerminateAPIRequest(v *AlibabaWdkCouponTemplateTerminateAPIRequest) {
	v.Reset()
	poolAlibabaWdkCouponTemplateTerminateAPIRequest.Put(v)
}
