package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundDetailAPIRequest 商旅机票退票详情接口 API请求
// alitrip.btrip.flight.distribution.refund.detail
//
// 商旅机票分销退票详情
type AlitripBtripFlightDistributionRefundDetailAPIRequest struct {
	model.Params
	// 获取退票单详情入参
	_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq
}

// NewAlitripBtripFlightDistributionRefundDetailRequest 初始化AlitripBtripFlightDistributionRefundDetailAPIRequest对象
func NewAlitripBtripFlightDistributionRefundDetailRequest() *AlitripBtripFlightDistributionRefundDetailAPIRequest {
	return &AlitripBtripFlightDistributionRefundDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionRefundDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionRefundDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionRefundDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 获取退票单详情入参
func (r *AlitripBtripFlightDistributionRefundDetailAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripBtripFlightDistributionRefundDetailAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightOrderOperateCommonRq {
	return r._paramBtripFlightOrderOperateCommonRq
}
