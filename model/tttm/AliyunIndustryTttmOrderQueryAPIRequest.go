package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunindustrytttmorderqueryAPIRequest 天天特卖数字工厂订单获取 API请求
// aliyun.industry.tttm.order.query
//
// 获取阿里云数字工厂内天天特卖业务的订单
type AliyunindustrytttmorderqueryAPIRequest struct {
	model.Params
	// 订单号
	_orderId string
	// 外部采购单号
	_externalId string
}

// NewAliyunindustrytttmorderqueryRequest 初始化AliyunindustrytttmorderqueryAPIRequest对象
func NewAliyunindustrytttmorderqueryRequest() *AliyunindustrytttmorderqueryAPIRequest {
	return &AliyunindustrytttmorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunindustrytttmorderqueryAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunindustrytttmorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunindustrytttmorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *AliyunindustrytttmorderqueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AliyunindustrytttmorderqueryAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetExternalId is ExternalId Setter
// 外部采购单号
func (r *AliyunindustrytttmorderqueryAPIRequest) SetExternalId(_externalId string) error {
	r._externalId = _externalId
	r.Set("external_id", _externalId)
	return nil
}

// GetExternalId ExternalId Getter
func (r AliyunindustrytttmorderqueryAPIRequest) GetExternalId() string {
	return r._externalId
}
