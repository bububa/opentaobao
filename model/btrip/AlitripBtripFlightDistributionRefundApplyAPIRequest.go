package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundApplyAPIRequest 商旅机票分销-退票申请 API请求
// alitrip.btrip.flight.distribution.refund.apply
//
// 商旅机票分销-退票申请
type AlitripBtripFlightDistributionRefundApplyAPIRequest struct {
	model.Params
	// 退票申请入参
	_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq
}

// NewAlitripBtripFlightDistributionRefundApplyRequest 初始化AlitripBtripFlightDistributionRefundApplyAPIRequest对象
func NewAlitripBtripFlightDistributionRefundApplyRequest() *AlitripBtripFlightDistributionRefundApplyAPIRequest {
	return &AlitripBtripFlightDistributionRefundApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionRefundApplyAPIRequest) Reset() {
	r._paramBtripFlightRefundApplyRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionRefundApplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionRefundApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionRefundApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightRefundApplyRq is ParamBtripFlightRefundApplyRq Setter
// 退票申请入参
func (r *AlitripBtripFlightDistributionRefundApplyAPIRequest) SetParamBtripFlightRefundApplyRq(_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq) error {
	r._paramBtripFlightRefundApplyRq = _paramBtripFlightRefundApplyRq
	r.Set("param_btrip_flight_refund_apply_rq", _paramBtripFlightRefundApplyRq)
	return nil
}

// GetParamBtripFlightRefundApplyRq ParamBtripFlightRefundApplyRq Getter
func (r AlitripBtripFlightDistributionRefundApplyAPIRequest) GetParamBtripFlightRefundApplyRq() *BtripFlightRefundApplyRq {
	return r._paramBtripFlightRefundApplyRq
}

var poolAlitripBtripFlightDistributionRefundApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionRefundApplyRequest()
	},
}

// GetAlitripBtripFlightDistributionRefundApplyRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionRefundApplyAPIRequest
func GetAlitripBtripFlightDistributionRefundApplyAPIRequest() *AlitripBtripFlightDistributionRefundApplyAPIRequest {
	return poolAlitripBtripFlightDistributionRefundApplyAPIRequest.Get().(*AlitripBtripFlightDistributionRefundApplyAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionRefundApplyAPIRequest 将 AlitripBtripFlightDistributionRefundApplyAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionRefundApplyAPIRequest(v *AlitripBtripFlightDistributionRefundApplyAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionRefundApplyAPIRequest.Put(v)
}
