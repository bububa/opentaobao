package idleisv

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvOrderAdjustpriceAPIRequest) Reset() {
	r._paramAdjustOrderPrice = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderAdjustpriceAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.adjustprice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderAdjustpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvOrderAdjustpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIdleIsvOrderAdjustpriceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvOrderAdjustpriceRequest()
	},
}

// GetAlibabaIdleIsvOrderAdjustpriceRequest 从 sync.Pool 获取 AlibabaIdleIsvOrderAdjustpriceAPIRequest
func GetAlibabaIdleIsvOrderAdjustpriceAPIRequest() *AlibabaIdleIsvOrderAdjustpriceAPIRequest {
	return poolAlibabaIdleIsvOrderAdjustpriceAPIRequest.Get().(*AlibabaIdleIsvOrderAdjustpriceAPIRequest)
}

// ReleaseAlibabaIdleIsvOrderAdjustpriceAPIRequest 将 AlibabaIdleIsvOrderAdjustpriceAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvOrderAdjustpriceAPIRequest(v *AlibabaIdleIsvOrderAdjustpriceAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvOrderAdjustpriceAPIRequest.Put(v)
}
