package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketetickettimeexpandAPIRequest 订单延时接口 API请求
// taobao.vmarket.eticket.time.expand
//
// 提供码商操作订单延期接口
type TaobaovmarketetickettimeexpandAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
	// 延长天数，延长时间=当前过期时间+延长天数
	_expandDays int64
}

// NewTaobaovmarketetickettimeexpandRequest 初始化TaobaovmarketetickettimeexpandAPIRequest对象
func NewTaobaovmarketetickettimeexpandRequest() *TaobaovmarketetickettimeexpandAPIRequest {
	return &TaobaovmarketetickettimeexpandAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovmarketetickettimeexpandAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.time.expand"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovmarketetickettimeexpandAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovmarketetickettimeexpandAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaovmarketetickettimeexpandAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaovmarketetickettimeexpandAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetExpandDays is ExpandDays Setter
// 延长天数，延长时间=当前过期时间+延长天数
func (r *TaobaovmarketetickettimeexpandAPIRequest) SetExpandDays(_expandDays int64) error {
	r._expandDays = _expandDays
	r.Set("expand_days", _expandDays)
	return nil
}

// GetExpandDays ExpandDays Getter
func (r TaobaovmarketetickettimeexpandAPIRequest) GetExpandDays() int64 {
	return r._expandDays
}
