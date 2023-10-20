package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewcancelAPIRequest 商旅机票改签取消 API请求
// alitrip.btrip.flight.distribution.change.newcancel
//
// 商旅机票改签取消
type AlitripBtripFlightDistributionChangeNewcancelAPIRequest struct {
	model.Params
	// 改签取消输入参数
	_paramBtripFlightModifyCancelRq *BtripFlightModifyCancelRq
}

// NewAlitripBtripFlightDistributionChangeNewcancelRequest 初始化AlitripBtripFlightDistributionChangeNewcancelAPIRequest对象
func NewAlitripBtripFlightDistributionChangeNewcancelRequest() *AlitripBtripFlightDistributionChangeNewcancelAPIRequest {
	return &AlitripBtripFlightDistributionChangeNewcancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionChangeNewcancelAPIRequest) Reset() {
	r._paramBtripFlightModifyCancelRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeNewcancelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newcancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeNewcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionChangeNewcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyCancelRq is ParamBtripFlightModifyCancelRq Setter
// 改签取消输入参数
func (r *AlitripBtripFlightDistributionChangeNewcancelAPIRequest) SetParamBtripFlightModifyCancelRq(_paramBtripFlightModifyCancelRq *BtripFlightModifyCancelRq) error {
	r._paramBtripFlightModifyCancelRq = _paramBtripFlightModifyCancelRq
	r.Set("param_btrip_flight_modify_cancel_rq", _paramBtripFlightModifyCancelRq)
	return nil
}

// GetParamBtripFlightModifyCancelRq ParamBtripFlightModifyCancelRq Getter
func (r AlitripBtripFlightDistributionChangeNewcancelAPIRequest) GetParamBtripFlightModifyCancelRq() *BtripFlightModifyCancelRq {
	return r._paramBtripFlightModifyCancelRq
}

var poolAlitripBtripFlightDistributionChangeNewcancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionChangeNewcancelRequest()
	},
}

// GetAlitripBtripFlightDistributionChangeNewcancelRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeNewcancelAPIRequest
func GetAlitripBtripFlightDistributionChangeNewcancelAPIRequest() *AlitripBtripFlightDistributionChangeNewcancelAPIRequest {
	return poolAlitripBtripFlightDistributionChangeNewcancelAPIRequest.Get().(*AlitripBtripFlightDistributionChangeNewcancelAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionChangeNewcancelAPIRequest 将 AlitripBtripFlightDistributionChangeNewcancelAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeNewcancelAPIRequest(v *AlitripBtripFlightDistributionChangeNewcancelAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeNewcancelAPIRequest.Put(v)
}
