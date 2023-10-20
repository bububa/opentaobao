package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponSpreadApplyAPIRequest 普通发券 API请求
// alibaba.wdk.coupon.spread.apply
//
// 优惠券发放
type AlibabaWdkCouponSpreadApplyAPIRequest struct {
	model.Params
	// 参数对象
	_paramWdkCouponApplyParam *WdkCouponApplyParam
}

// NewAlibabaWdkCouponSpreadApplyRequest 初始化AlibabaWdkCouponSpreadApplyAPIRequest对象
func NewAlibabaWdkCouponSpreadApplyRequest() *AlibabaWdkCouponSpreadApplyAPIRequest {
	return &AlibabaWdkCouponSpreadApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkCouponSpreadApplyAPIRequest) Reset() {
	r._paramWdkCouponApplyParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSpreadApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.spread.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSpreadApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkCouponSpreadApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWdkCouponApplyParam is ParamWdkCouponApplyParam Setter
// 参数对象
func (r *AlibabaWdkCouponSpreadApplyAPIRequest) SetParamWdkCouponApplyParam(_paramWdkCouponApplyParam *WdkCouponApplyParam) error {
	r._paramWdkCouponApplyParam = _paramWdkCouponApplyParam
	r.Set("param_wdk_coupon_apply_param", _paramWdkCouponApplyParam)
	return nil
}

// GetParamWdkCouponApplyParam ParamWdkCouponApplyParam Getter
func (r AlibabaWdkCouponSpreadApplyAPIRequest) GetParamWdkCouponApplyParam() *WdkCouponApplyParam {
	return r._paramWdkCouponApplyParam
}

var poolAlibabaWdkCouponSpreadApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkCouponSpreadApplyRequest()
	},
}

// GetAlibabaWdkCouponSpreadApplyRequest 从 sync.Pool 获取 AlibabaWdkCouponSpreadApplyAPIRequest
func GetAlibabaWdkCouponSpreadApplyAPIRequest() *AlibabaWdkCouponSpreadApplyAPIRequest {
	return poolAlibabaWdkCouponSpreadApplyAPIRequest.Get().(*AlibabaWdkCouponSpreadApplyAPIRequest)
}

// ReleaseAlibabaWdkCouponSpreadApplyAPIRequest 将 AlibabaWdkCouponSpreadApplyAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkCouponSpreadApplyAPIRequest(v *AlibabaWdkCouponSpreadApplyAPIRequest) {
	v.Reset()
	poolAlibabaWdkCouponSpreadApplyAPIRequest.Put(v)
}
