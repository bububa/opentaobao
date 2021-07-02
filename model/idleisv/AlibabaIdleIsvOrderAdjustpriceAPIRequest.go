package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvOrderAdjustpriceAPIRequest 闲鱼服务商订单价格修改接口 API请求
// alibaba.idle.isv.order.adjustprice
//
// 闲鱼用户通过授权的服务商修改订单价格和邮费
type AlibabaIdleIsvOrderAdjustpriceAPIRequest struct {
	model.Params
	// 输入参数
	_paramAdjustOrderPrice *IsvAdjustOrderPriceDto
}

// NewAlibabaIdleIsvOrderAdjustpriceRequest 初始化AlibabaIdleIsvOrderAdjustpriceAPIRequest对象
func NewAlibabaIdleIsvOrderAdjustpriceRequest() *AlibabaIdleIsvOrderAdjustpriceAPIRequest {
	return &AlibabaIdleIsvOrderAdjustpriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderAdjustpriceAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.adjustprice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderAdjustpriceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamAdjustOrderPrice is ParamAdjustOrderPrice Setter
// 输入参数
func (r *AlibabaIdleIsvOrderAdjustpriceAPIRequest) SetParamAdjustOrderPrice(_paramAdjustOrderPrice *IsvAdjustOrderPriceDto) error {
	r._paramAdjustOrderPrice = _paramAdjustOrderPrice
	r.Set("param_adjust_order_price", _paramAdjustOrderPrice)
	return nil
}

// GetParamAdjustOrderPrice ParamAdjustOrderPrice Getter
func (r AlibabaIdleIsvOrderAdjustpriceAPIRequest) GetParamAdjustOrderPrice() *IsvAdjustOrderPriceDto {
	return r._paramAdjustOrderPrice
}
