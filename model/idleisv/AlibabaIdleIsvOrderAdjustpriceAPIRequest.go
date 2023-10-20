package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvorderadjustpriceAPIRequest 闲鱼服务商订单价格修改接口 API请求
// alibaba.idle.isv.order.adjustprice
//
// 闲鱼用户通过授权的服务商修改订单价格和邮费
type AlibabaidleisvorderadjustpriceAPIRequest struct {
	model.Params
	// 输入参数
	_paramAdjustOrderPrice *IsvAdjustOrderPriceDto
}

// NewAlibabaidleisvorderadjustpriceRequest 初始化AlibabaidleisvorderadjustpriceAPIRequest对象
func NewAlibabaidleisvorderadjustpriceRequest() *AlibabaidleisvorderadjustpriceAPIRequest {
	return &AlibabaidleisvorderadjustpriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvorderadjustpriceAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.adjustprice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvorderadjustpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvorderadjustpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAdjustOrderPrice is ParamAdjustOrderPrice Setter
// 输入参数
func (r *AlibabaidleisvorderadjustpriceAPIRequest) SetParamAdjustOrderPrice(_paramAdjustOrderPrice *IsvAdjustOrderPriceDto) error {
	r._paramAdjustOrderPrice = _paramAdjustOrderPrice
	r.Set("param_adjust_order_price", _paramAdjustOrderPrice)
	return nil
}

// GetParamAdjustOrderPrice ParamAdjustOrderPrice Getter
func (r AlibabaidleisvorderadjustpriceAPIRequest) GetParamAdjustOrderPrice() *IsvAdjustOrderPriceDto {
	return r._paramAdjustOrderPrice
}
