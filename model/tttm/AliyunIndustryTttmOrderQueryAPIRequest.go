package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunIndustryTttmOrderQueryAPIRequest 天天特卖数字工厂订单获取 API请求
// aliyun.industry.tttm.order.query
//
// 获取阿里云数字工厂内天天特卖业务的订单
type AliyunIndustryTttmOrderQueryAPIRequest struct {
	model.Params
	// 订单号
	_orderId string
	// 外部采购单号
	_externalId string
}

// NewAliyunIndustryTttmOrderQueryRequest 初始化AliyunIndustryTttmOrderQueryAPIRequest对象
func NewAliyunIndustryTttmOrderQueryRequest() *AliyunIndustryTttmOrderQueryAPIRequest {
	return &AliyunIndustryTttmOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmOrderQueryAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *AliyunIndustryTttmOrderQueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AliyunIndustryTttmOrderQueryAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetExternalId is ExternalId Setter
// 外部采购单号
func (r *AliyunIndustryTttmOrderQueryAPIRequest) SetExternalId(_externalId string) error {
	r._externalId = _externalId
	r.Set("external_id", _externalId)
	return nil
}

// GetExternalId ExternalId Getter
func (r AliyunIndustryTttmOrderQueryAPIRequest) GetExternalId() string {
	return r._externalId
}
