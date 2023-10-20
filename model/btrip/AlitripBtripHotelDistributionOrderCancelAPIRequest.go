package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriphoteldistributionordercancelAPIRequest 商旅酒店API分销取消订单 API请求
// alitrip.btrip.hotel.distribution.order.cancel
//
// 商旅酒店API分销取消订单
type AlitripbtriphoteldistributionordercancelAPIRequest struct {
	model.Params
	// 取消订单接口入参
	_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq
}

// NewAlitripbtriphoteldistributionordercancelRequest 初始化AlitripbtriphoteldistributionordercancelAPIRequest对象
func NewAlitripbtriphoteldistributionordercancelRequest() *AlitripbtriphoteldistributionordercancelAPIRequest {
	return &AlitripbtriphoteldistributionordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtriphoteldistributionordercancelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtriphoteldistributionordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtriphoteldistributionordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripHotelOrderOperateRq is ParamBtripHotelOrderOperateRq Setter
// 取消订单接口入参
func (r *AlitripbtriphoteldistributionordercancelAPIRequest) SetParamBtripHotelOrderOperateRq(_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
	r._paramBtripHotelOrderOperateRq = _paramBtripHotelOrderOperateRq
	r.Set("param_btrip_hotel_order_operate_rq", _paramBtripHotelOrderOperateRq)
	return nil
}

// GetParamBtripHotelOrderOperateRq ParamBtripHotelOrderOperateRq Getter
func (r AlitripbtriphoteldistributionordercancelAPIRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
	return r._paramBtripHotelOrderOperateRq
}
