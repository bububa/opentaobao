package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponAbandonAPIRequest 废券 API请求
// alibaba.wdk.coupon.abandon
//
// 优惠券废弃
type AlibabaWdkCouponAbandonAPIRequest struct {
	model.Params
	// 废券参数
	_paramWdkCouponAbandonParam *WdkCouponAbandonParam
}

// NewAlibabaWdkCouponAbandonRequest 初始化AlibabaWdkCouponAbandonAPIRequest对象
func NewAlibabaWdkCouponAbandonRequest() *AlibabaWdkCouponAbandonAPIRequest {
	return &AlibabaWdkCouponAbandonAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkCouponAbandonAPIRequest) Reset() {
	r._paramWdkCouponAbandonParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponAbandonAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.abandon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponAbandonAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkCouponAbandonAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWdkCouponAbandonParam is ParamWdkCouponAbandonParam Setter
// 废券参数
func (r *AlibabaWdkCouponAbandonAPIRequest) SetParamWdkCouponAbandonParam(_paramWdkCouponAbandonParam *WdkCouponAbandonParam) error {
	r._paramWdkCouponAbandonParam = _paramWdkCouponAbandonParam
	r.Set("param_wdk_coupon_abandon_param", _paramWdkCouponAbandonParam)
	return nil
}

// GetParamWdkCouponAbandonParam ParamWdkCouponAbandonParam Getter
func (r AlibabaWdkCouponAbandonAPIRequest) GetParamWdkCouponAbandonParam() *WdkCouponAbandonParam {
	return r._paramWdkCouponAbandonParam
}

var poolAlibabaWdkCouponAbandonAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkCouponAbandonRequest()
	},
}

// GetAlibabaWdkCouponAbandonRequest 从 sync.Pool 获取 AlibabaWdkCouponAbandonAPIRequest
func GetAlibabaWdkCouponAbandonAPIRequest() *AlibabaWdkCouponAbandonAPIRequest {
	return poolAlibabaWdkCouponAbandonAPIRequest.Get().(*AlibabaWdkCouponAbandonAPIRequest)
}

// ReleaseAlibabaWdkCouponAbandonAPIRequest 将 AlibabaWdkCouponAbandonAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkCouponAbandonAPIRequest(v *AlibabaWdkCouponAbandonAPIRequest) {
	v.Reset()
	poolAlibabaWdkCouponAbandonAPIRequest.Put(v)
}
