package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriphoteldistributionorderpayAPIRequest 商旅酒店分销订单支付 API请求
// alitrip.btrip.hotel.distribution.order.pay
//
// 商旅酒店分销订单支付
type AlitripbtriphoteldistributionorderpayAPIRequest struct {
	model.Params
	// 通知商旅支付成功接口参数
	_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq
}

// NewAlitripbtriphoteldistributionorderpayRequest 初始化AlitripbtriphoteldistributionorderpayAPIRequest对象
func NewAlitripbtriphoteldistributionorderpayRequest() *AlitripbtriphoteldistributionorderpayAPIRequest {
	return &AlitripbtriphoteldistributionorderpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtriphoteldistributionorderpayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtriphoteldistributionorderpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtriphoteldistributionorderpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripHotelOrderOperateRq is ParamBtripHotelOrderOperateRq Setter
// 通知商旅支付成功接口参数
func (r *AlitripbtriphoteldistributionorderpayAPIRequest) SetParamBtripHotelOrderOperateRq(_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
	r._paramBtripHotelOrderOperateRq = _paramBtripHotelOrderOperateRq
	r.Set("param_btrip_hotel_order_operate_rq", _paramBtripHotelOrderOperateRq)
	return nil
}

// GetParamBtripHotelOrderOperateRq ParamBtripHotelOrderOperateRq Getter
func (r AlitripbtriphoteldistributionorderpayAPIRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
	return r._paramBtripHotelOrderOperateRq
}
