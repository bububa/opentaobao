package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewapplyAPIRequest 商旅机票改签申请V2 API请求
// alitrip.btrip.flight.distribution.change.newapply
//
// 商旅机票改签申请
type AlitripBtripFlightDistributionChangeNewapplyAPIRequest struct {
	model.Params
	// 改签申请入参
	_paramBtripFlightModifyApplyRq *BtripFlightModifyApplyRq
}

// NewAlitripBtripFlightDistributionChangeNewapplyRequest 初始化AlitripBtripFlightDistributionChangeNewapplyAPIRequest对象
func NewAlitripBtripFlightDistributionChangeNewapplyRequest() *AlitripBtripFlightDistributionChangeNewapplyAPIRequest {
	return &AlitripBtripFlightDistributionChangeNewapplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionChangeNewapplyAPIRequest) Reset() {
	r._paramBtripFlightModifyApplyRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeNewapplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeNewapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionChangeNewapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyApplyRq is ParamBtripFlightModifyApplyRq Setter
// 改签申请入参
func (r *AlitripBtripFlightDistributionChangeNewapplyAPIRequest) SetParamBtripFlightModifyApplyRq(_paramBtripFlightModifyApplyRq *BtripFlightModifyApplyRq) error {
	r._paramBtripFlightModifyApplyRq = _paramBtripFlightModifyApplyRq
	r.Set("param_btrip_flight_modify_apply_rq", _paramBtripFlightModifyApplyRq)
	return nil
}

// GetParamBtripFlightModifyApplyRq ParamBtripFlightModifyApplyRq Getter
func (r AlitripBtripFlightDistributionChangeNewapplyAPIRequest) GetParamBtripFlightModifyApplyRq() *BtripFlightModifyApplyRq {
	return r._paramBtripFlightModifyApplyRq
}

var poolAlitripBtripFlightDistributionChangeNewapplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionChangeNewapplyRequest()
	},
}

// GetAlitripBtripFlightDistributionChangeNewapplyRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeNewapplyAPIRequest
func GetAlitripBtripFlightDistributionChangeNewapplyAPIRequest() *AlitripBtripFlightDistributionChangeNewapplyAPIRequest {
	return poolAlitripBtripFlightDistributionChangeNewapplyAPIRequest.Get().(*AlitripBtripFlightDistributionChangeNewapplyAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionChangeNewapplyAPIRequest 将 AlitripBtripFlightDistributionChangeNewapplyAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeNewapplyAPIRequest(v *AlitripBtripFlightDistributionChangeNewapplyAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeNewapplyAPIRequest.Put(v)
}
