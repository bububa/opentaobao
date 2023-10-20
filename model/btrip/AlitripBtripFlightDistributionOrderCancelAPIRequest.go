package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionordercancelAPIRequest 商旅机票分销-取消订单 API请求
// alitrip.btrip.flight.distribution.order.cancel
//
// 商旅机票分销取消订单
type AlitripbtripflightdistributionordercancelAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramBtripFlightOrderOperateCommonRq *BtripFlightCancelOrderRq
}

// NewAlitripbtripflightdistributionordercancelRequest 初始化AlitripbtripflightdistributionordercancelAPIRequest对象
func NewAlitripbtripflightdistributionordercancelRequest() *AlitripbtripflightdistributionordercancelAPIRequest {
	return &AlitripbtripflightdistributionordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionordercancelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 系统自动生成
func (r *AlitripbtripflightdistributionordercancelAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightCancelOrderRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripbtripflightdistributionordercancelAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightCancelOrderRq {
	return r._paramBtripFlightOrderOperateCommonRq
}
