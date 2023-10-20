package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripHotelDistributionOrderPayAPIRequest) Reset() {
	r._paramBtripHotelOrderOperateRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderPayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderPayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripHotelDistributionOrderPayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripHotelOrderOperateRq is ParamBtripHotelOrderOperateRq Setter
// 通知商旅支付成功接口参数
func (r *AlitripBtripHotelDistributionOrderPayAPIRequest) SetParamBtripHotelOrderOperateRq(_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq) error {
	r._paramBtripHotelOrderOperateRq = _paramBtripHotelOrderOperateRq
	r.Set("param_btrip_hotel_order_operate_rq", _paramBtripHotelOrderOperateRq)
	return nil
}

// GetParamBtripHotelOrderOperateRq ParamBtripHotelOrderOperateRq Getter
func (r AlitripBtripHotelDistributionOrderPayAPIRequest) GetParamBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
	return r._paramBtripHotelOrderOperateRq
}

var poolAlitripBtripHotelDistributionOrderPayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripHotelDistributionOrderPayRequest()
	},
}

// GetAlitripBtripHotelDistributionOrderPayRequest 从 sync.Pool 获取 AlitripBtripHotelDistributionOrderPayAPIRequest
func GetAlitripBtripHotelDistributionOrderPayAPIRequest() *AlitripBtripHotelDistributionOrderPayAPIRequest {
	return poolAlitripBtripHotelDistributionOrderPayAPIRequest.Get().(*AlitripBtripHotelDistributionOrderPayAPIRequest)
}

// ReleaseAlitripBtripHotelDistributionOrderPayAPIRequest 将 AlitripBtripHotelDistributionOrderPayAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripHotelDistributionOrderPayAPIRequest(v *AlitripBtripHotelDistributionOrderPayAPIRequest) {
	v.Reset()
	poolAlitripBtripHotelDistributionOrderPayAPIRequest.Put(v)
}
