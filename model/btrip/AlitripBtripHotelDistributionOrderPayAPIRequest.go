package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderPayAPIRequest 商旅酒店分销订单支付 API请求
// alitrip.btrip.hotel.distribution.order.pay
//
// 商旅酒店分销订单支付
type AlitripBtripHotelDistributionOrderPayAPIRequest struct {
	model.Params
	// 通知商旅支付成功接口参数
	_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq
}

// NewAlitripBtripHotelDistributionOrderPayRequest 初始化AlitripBtripHotelDistributionOrderPayAPIRequest对象
func NewAlitripBtripHotelDistributionOrderPayRequest() *AlitripBtripHotelDistributionOrderPayAPIRequest {
	return &AlitripBtripHotelDistributionOrderPayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderPayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderPayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamBtripHotelOrderOperateRq Setter
// 通知商旅支付成功接口参数
func (r *AlitripBtripHotelDistributionOrderPayAPIRequest) SetParamBtripHotelOrderOperateRq(_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
	r._paramBtripHotelOrderOperateRq = _paramBtripHotelOrderOperateRq
	r.Set("param_btrip_hotel_order_operate_rq", _paramBtripHotelOrderOperateRq)
	return nil
}

// Get ParamBtripHotelOrderOperateRq Getter
func (r AlitripBtripHotelDistributionOrderPayAPIRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
	return r._paramBtripHotelOrderOperateRq
}
