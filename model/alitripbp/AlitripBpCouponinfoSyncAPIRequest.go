package alitripbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBpCouponinfoSyncAPIRequest 飞猪广告券信息同步接口 API请求
// alitrip.bp.couponinfo.sync
//
// 飞猪商业化券信息同步
type AlitripBpCouponinfoSyncAPIRequest struct {
	model.Params
	// 商业化券同步接口请求
	_paramCouponDataRequest *CouponDataRequest
}

// NewAlitripBpCouponinfoSyncRequest 初始化AlitripBpCouponinfoSyncAPIRequest对象
func NewAlitripBpCouponinfoSyncRequest() *AlitripBpCouponinfoSyncAPIRequest {
	return &AlitripBpCouponinfoSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBpCouponinfoSyncAPIRequest) GetApiMethodName() string {
	return "alitrip.bp.couponinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBpCouponinfoSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBpCouponinfoSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponDataRequest is ParamCouponDataRequest Setter
// 商业化券同步接口请求
func (r *AlitripBpCouponinfoSyncAPIRequest) SetParamCouponDataRequest(_paramCouponDataRequest *CouponDataRequest) error {
	r._paramCouponDataRequest = _paramCouponDataRequest
	r.Set("param_coupon_data_request", _paramCouponDataRequest)
	return nil
}

// GetParamCouponDataRequest ParamCouponDataRequest Getter
func (r AlitripBpCouponinfoSyncAPIRequest) GetParamCouponDataRequest() *CouponDataRequest {
	return r._paramCouponDataRequest
}
