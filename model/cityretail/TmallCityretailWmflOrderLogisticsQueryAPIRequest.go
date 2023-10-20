package cityretail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcityretailwmflorderlogisticsqueryAPIRequest 完美履约订单物流状态查询接口 API请求
// tmall.cityretail.wmfl.order.logistics.query
//
// 完美履约订单物流状态查询接口，该接口只能查询未完结的履约单以及完结的3天内订单
type TmallcityretailwmflorderlogisticsqueryAPIRequest struct {
	model.Params
	// 订单号
	_mainOrderId string
}

// NewTmallcityretailwmflorderlogisticsqueryRequest 初始化TmallcityretailwmflorderlogisticsqueryAPIRequest对象
func NewTmallcityretailwmflorderlogisticsqueryRequest() *TmallcityretailwmflorderlogisticsqueryAPIRequest {
	return &TmallcityretailwmflorderlogisticsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcityretailwmflorderlogisticsqueryAPIRequest) GetApiMethodName() string {
	return "tmall.cityretail.wmfl.order.logistics.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcityretailwmflorderlogisticsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcityretailwmflorderlogisticsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 订单号
func (r *TmallcityretailwmflorderlogisticsqueryAPIRequest) SetMainOrderId(_mainOrderId string) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TmallcityretailwmflorderlogisticsqueryAPIRequest) GetMainOrderId() string {
	return r._mainOrderId
}
