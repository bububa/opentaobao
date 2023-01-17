package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundNewdetailAPIRequest 商旅机票退票详情接口V2 API请求
// alitrip.btrip.flight.distribution.refund.newdetail
//
// 商旅机票退票详情接口V2
type AlitripBtripFlightDistributionRefundNewdetailAPIRequest struct {
	model.Params
	// 获取退票单详情入参
	_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq
}

// NewAlitripBtripFlightDistributionRefundNewdetailRequest 初始化AlitripBtripFlightDistributionRefundNewdetailAPIRequest对象
func NewAlitripBtripFlightDistributionRefundNewdetailRequest() *AlitripBtripFlightDistributionRefundNewdetailAPIRequest {
	return &AlitripBtripFlightDistributionRefundNewdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionRefundNewdetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.newdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionRefundNewdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionRefundNewdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightOrderOperateCommonRq is ParamBtripFlightOrderOperateCommonRq Setter
// 获取退票单详情入参
func (r *AlitripBtripFlightDistributionRefundNewdetailAPIRequest) SetParamBtripFlightOrderOperateCommonRq(_paramBtripFlightOrderOperateCommonRq *BtripFlightOrderOperateCommonRq) error {
	r._paramBtripFlightOrderOperateCommonRq = _paramBtripFlightOrderOperateCommonRq
	r.Set("param_btrip_flight_order_operate_common_rq", _paramBtripFlightOrderOperateCommonRq)
	return nil
}

// GetParamBtripFlightOrderOperateCommonRq ParamBtripFlightOrderOperateCommonRq Getter
func (r AlitripBtripFlightDistributionRefundNewdetailAPIRequest) GetParamBtripFlightOrderOperateCommonRq() *BtripFlightOrderOperateCommonRq {
	return r._paramBtripFlightOrderOperateCommonRq
}
