package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeApplyAPIRequest 商旅机票改签申请 API请求
// alitrip.btrip.flight.distribution.change.apply
//
// 商旅机票改签申请
type AlitripBtripFlightDistributionChangeApplyAPIRequest struct {
	model.Params
	// 改签申请入参
	_paramBtripFlightModifyApplyRq *BtripFlightModifyApplyRq
}

// NewAlitripBtripFlightDistributionChangeApplyRequest 初始化AlitripBtripFlightDistributionChangeApplyAPIRequest对象
func NewAlitripBtripFlightDistributionChangeApplyRequest() *AlitripBtripFlightDistributionChangeApplyAPIRequest {
	return &AlitripBtripFlightDistributionChangeApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionChangeApplyAPIRequest) Reset() {
	r._paramBtripFlightModifyApplyRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeApplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionChangeApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyApplyRq is ParamBtripFlightModifyApplyRq Setter
// 改签申请入参
func (r *AlitripBtripFlightDistributionChangeApplyAPIRequest) SetParamBtripFlightModifyApplyRq(_paramBtripFlightModifyApplyRq *BtripFlightModifyApplyRq) error {
	r._paramBtripFlightModifyApplyRq = _paramBtripFlightModifyApplyRq
	r.Set("param_btrip_flight_modify_apply_rq", _paramBtripFlightModifyApplyRq)
	return nil
}

// GetParamBtripFlightModifyApplyRq ParamBtripFlightModifyApplyRq Getter
func (r AlitripBtripFlightDistributionChangeApplyAPIRequest) GetParamBtripFlightModifyApplyRq() *BtripFlightModifyApplyRq {
	return r._paramBtripFlightModifyApplyRq
}

var poolAlitripBtripFlightDistributionChangeApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionChangeApplyRequest()
	},
}

// GetAlitripBtripFlightDistributionChangeApplyRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeApplyAPIRequest
func GetAlitripBtripFlightDistributionChangeApplyAPIRequest() *AlitripBtripFlightDistributionChangeApplyAPIRequest {
	return poolAlitripBtripFlightDistributionChangeApplyAPIRequest.Get().(*AlitripBtripFlightDistributionChangeApplyAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionChangeApplyAPIRequest 将 AlitripBtripFlightDistributionChangeApplyAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeApplyAPIRequest(v *AlitripBtripFlightDistributionChangeApplyAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeApplyAPIRequest.Put(v)
}
