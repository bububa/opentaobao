package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewdetailAPIRequest 商旅机票改签详情接口 API请求
// alitrip.btrip.flight.distribution.change.newdetail
//
// 商旅机票改签详情接口
type AlitripBtripFlightDistributionChangeNewdetailAPIRequest struct {
	model.Params
	// 获取改签单详情入参
	_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq
}

// NewAlitripBtripFlightDistributionChangeNewdetailRequest 初始化AlitripBtripFlightDistributionChangeNewdetailAPIRequest对象
func NewAlitripBtripFlightDistributionChangeNewdetailRequest() *AlitripBtripFlightDistributionChangeNewdetailAPIRequest {
	return &AlitripBtripFlightDistributionChangeNewdetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionChangeNewdetailAPIRequest) Reset() {
	r._paramBtripFlightOrderOperateCommonRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeNewdetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeNewdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionChangeNewdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 获取改签单详情入参
func (r *AlitripBtripFlightDistributionChangeNewdetailAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripBtripFlightDistributionChangeNewdetailAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightOrderOperateCommonRq {
	return r._paramBtripFlightOrderOperateCommonRq
}

var poolAlitripBtripFlightDistributionChangeNewdetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionChangeNewdetailRequest()
	},
}

// GetAlitripBtripFlightDistributionChangeNewdetailRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeNewdetailAPIRequest
func GetAlitripBtripFlightDistributionChangeNewdetailAPIRequest() *AlitripBtripFlightDistributionChangeNewdetailAPIRequest {
	return poolAlitripBtripFlightDistributionChangeNewdetailAPIRequest.Get().(*AlitripBtripFlightDistributionChangeNewdetailAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionChangeNewdetailAPIRequest 将 AlitripBtripFlightDistributionChangeNewdetailAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeNewdetailAPIRequest(v *AlitripBtripFlightDistributionChangeNewdetailAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeNewdetailAPIRequest.Put(v)
}
