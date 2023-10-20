package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangenewdetailAPIRequest 商旅机票改签详情接口 API请求
// alitrip.btrip.flight.distribution.change.newdetail
//
// 商旅机票改签详情接口
type AlitripbtripflightdistributionchangenewdetailAPIRequest struct {
	model.Params
	// 获取改签单详情入参
	_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq
}

// NewAlitripbtripflightdistributionchangenewdetailRequest 初始化AlitripbtripflightdistributionchangenewdetailAPIRequest对象
func NewAlitripbtripflightdistributionchangenewdetailRequest() *AlitripbtripflightdistributionchangenewdetailAPIRequest {
	return &AlitripbtripflightdistributionchangenewdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionchangenewdetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionchangenewdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionchangenewdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 获取改签单详情入参
func (r *AlitripbtripflightdistributionchangenewdetailAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripbtripflightdistributionchangenewdetailAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightOrderOperateCommonRq {
	return r._paramBtripFlightOrderOperateCommonRq
}
