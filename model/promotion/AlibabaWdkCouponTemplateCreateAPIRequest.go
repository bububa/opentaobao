package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateCreateAPIRequest 优惠券模版创建 API请求
// alibaba.wdk.coupon.template.create
//
// 开放给外部商家创建优惠券模版
type AlibabaWdkCouponTemplateCreateAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest
}

// NewAlibabaWdkCouponTemplateCreateRequest 初始化AlibabaWdkCouponTemplateCreateAPIRequest对象
func NewAlibabaWdkCouponTemplateCreateRequest() *AlibabaWdkCouponTemplateCreateAPIRequest {
	return &AlibabaWdkCouponTemplateCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkCouponTemplateCreateAPIRequest) Reset() {
	r._paramCouponTemplateOperateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkCouponTemplateCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateOperateRequest is ParamCouponTemplateOperateRequest Setter
// 请求
func (r *AlibabaWdkCouponTemplateCreateAPIRequest) SetParamCouponTemplateOperateRequest(_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest) error {
	r._paramCouponTemplateOperateRequest = _paramCouponTemplateOperateRequest
	r.Set("param_coupon_template_operate_request", _paramCouponTemplateOperateRequest)
	return nil
}

// GetParamCouponTemplateOperateRequest ParamCouponTemplateOperateRequest Getter
func (r AlibabaWdkCouponTemplateCreateAPIRequest) GetParamCouponTemplateOperateRequest() *CouponTemplateOperateRequest {
	return r._paramCouponTemplateOperateRequest
}

var poolAlibabaWdkCouponTemplateCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkCouponTemplateCreateRequest()
	},
}

// GetAlibabaWdkCouponTemplateCreateRequest 从 sync.Pool 获取 AlibabaWdkCouponTemplateCreateAPIRequest
func GetAlibabaWdkCouponTemplateCreateAPIRequest() *AlibabaWdkCouponTemplateCreateAPIRequest {
	return poolAlibabaWdkCouponTemplateCreateAPIRequest.Get().(*AlibabaWdkCouponTemplateCreateAPIRequest)
}

// ReleaseAlibabaWdkCouponTemplateCreateAPIRequest 将 AlibabaWdkCouponTemplateCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkCouponTemplateCreateAPIRequest(v *AlibabaWdkCouponTemplateCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkCouponTemplateCreateAPIRequest.Put(v)
}
