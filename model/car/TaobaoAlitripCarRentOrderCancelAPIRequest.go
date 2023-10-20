package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripcarrentordercancelAPIRequest 租车-取消订单 API请求
// taobao.alitrip.car.rent.order.cancel
//
// 服务商主动取消用户订单或者拒绝取消订单.
type TaobaoalitripcarrentordercancelAPIRequest struct {
	model.Params
	// 取消请求对象
	_param0 *RentProviderCancelRequest
}

// NewTaobaoalitripcarrentordercancelRequest 初始化TaobaoalitripcarrentordercancelAPIRequest对象
func NewTaobaoalitripcarrentordercancelRequest() *TaobaoalitripcarrentordercancelAPIRequest {
	return &TaobaoalitripcarrentordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripcarrentordercancelAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.rent.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripcarrentordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripcarrentordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 取消请求对象
func (r *TaobaoalitripcarrentordercancelAPIRequest) SetParam0(_param0 *RentProviderCancelRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoalitripcarrentordercancelAPIRequest) GetParam0() *RentProviderCancelRequest {
	return r._param0
}
