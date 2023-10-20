package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateUpdateAPIRequest 优惠券模版修改 API请求
// alibaba.wdk.coupon.template.update
//
// 优惠券模版修改
type AlibabaWdkCouponTemplateUpdateAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest
}

// NewAlibabaWdkCouponTemplateUpdateRequest 初始化AlibabaWdkCouponTemplateUpdateAPIRequest对象
func NewAlibabaWdkCouponTemplateUpdateRequest() *AlibabaWdkCouponTemplateUpdateAPIRequest {
	return &AlibabaWdkCouponTemplateUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkCouponTemplateUpdateAPIRequest) Reset() {
	r._paramCouponTemplateOperateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkCouponTemplateUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateOperateRequest is ParamCouponTemplateOperateRequest Setter
// 请求
func (r *AlibabaWdkCouponTemplateUpdateAPIRequest) SetParamCouponTemplateOperateRequest(_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest) error {
	r._paramCouponTemplateOperateRequest = _paramCouponTemplateOperateRequest
	r.Set("param_coupon_template_operate_request", _paramCouponTemplateOperateRequest)
	return nil
}

// GetParamCouponTemplateOperateRequest ParamCouponTemplateOperateRequest Getter
func (r AlibabaWdkCouponTemplateUpdateAPIRequest) GetParamCouponTemplateOperateRequest() *CouponTemplateOperateRequest {
	return r._paramCouponTemplateOperateRequest
}

var poolAlibabaWdkCouponTemplateUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkCouponTemplateUpdateRequest()
	},
}

// GetAlibabaWdkCouponTemplateUpdateRequest 从 sync.Pool 获取 AlibabaWdkCouponTemplateUpdateAPIRequest
func GetAlibabaWdkCouponTemplateUpdateAPIRequest() *AlibabaWdkCouponTemplateUpdateAPIRequest {
	return poolAlibabaWdkCouponTemplateUpdateAPIRequest.Get().(*AlibabaWdkCouponTemplateUpdateAPIRequest)
}

// ReleaseAlibabaWdkCouponTemplateUpdateAPIRequest 将 AlibabaWdkCouponTemplateUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkCouponTemplateUpdateAPIRequest(v *AlibabaWdkCouponTemplateUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkCouponTemplateUpdateAPIRequest.Put(v)
}
