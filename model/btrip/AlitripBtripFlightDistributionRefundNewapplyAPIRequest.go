package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundNewapplyAPIRequest 商旅机票分销-退票申请 API请求
// alitrip.btrip.flight.distribution.refund.newapply
//
// 商旅机票分销-退票申请
type AlitripBtripFlightDistributionRefundNewapplyAPIRequest struct {
	model.Params
	// 退票申请入参
	_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq
}

// NewAlitripBtripFlightDistributionRefundNewapplyRequest 初始化AlitripBtripFlightDistributionRefundNewapplyAPIRequest对象
func NewAlitripBtripFlightDistributionRefundNewapplyRequest() *AlitripBtripFlightDistributionRefundNewapplyAPIRequest {
	return &AlitripBtripFlightDistributionRefundNewapplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionRefundNewapplyAPIRequest) Reset() {
	r._paramBtripFlightRefundApplyRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionRefundNewapplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.newapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionRefundNewapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionRefundNewapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightRefundApplyRq is ParamBtripFlightRefundApplyRq Setter
// 退票申请入参
func (r *AlitripBtripFlightDistributionRefundNewapplyAPIRequest) SetParamBtripFlightRefundApplyRq(_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq) error {
	r._paramBtripFlightRefundApplyRq = _paramBtripFlightRefundApplyRq
	r.Set("param_btrip_flight_refund_apply_rq", _paramBtripFlightRefundApplyRq)
	return nil
}

// GetParamBtripFlightRefundApplyRq ParamBtripFlightRefundApplyRq Getter
func (r AlitripBtripFlightDistributionRefundNewapplyAPIRequest) GetParamBtripFlightRefundApplyRq() *BtripFlightRefundApplyRq {
	return r._paramBtripFlightRefundApplyRq
}

var poolAlitripBtripFlightDistributionRefundNewapplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionRefundNewapplyRequest()
	},
}

// GetAlitripBtripFlightDistributionRefundNewapplyRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionRefundNewapplyAPIRequest
func GetAlitripBtripFlightDistributionRefundNewapplyAPIRequest() *AlitripBtripFlightDistributionRefundNewapplyAPIRequest {
	return poolAlitripBtripFlightDistributionRefundNewapplyAPIRequest.Get().(*AlitripBtripFlightDistributionRefundNewapplyAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionRefundNewapplyAPIRequest 将 AlitripBtripFlightDistributionRefundNewapplyAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionRefundNewapplyAPIRequest(v *AlitripBtripFlightDistributionRefundNewapplyAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionRefundNewapplyAPIRequest.Put(v)
}
