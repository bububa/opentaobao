package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkscmlrporderpredictAPIRequest 单量预测 API请求
// alibaba.wdk.scm.lrp.order.predict
//
// 提供基于门店和配送站的订单量预测，可用于提前安排人力资源
type AlibabawdkscmlrporderpredictAPIRequest struct {
	model.Params
	// 单量预测查询参数
	_paramOrderPredictQuery *OrderPredictQuery
}

// NewAlibabawdkscmlrporderpredictRequest 初始化AlibabawdkscmlrporderpredictAPIRequest对象
func NewAlibabawdkscmlrporderpredictRequest() *AlibabawdkscmlrporderpredictAPIRequest {
	return &AlibabawdkscmlrporderpredictAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkscmlrporderpredictAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.scm.lrp.order.predict"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkscmlrporderpredictAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkscmlrporderpredictAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderPredictQuery is ParamOrderPredictQuery Setter
// 单量预测查询参数
func (r *AlibabawdkscmlrporderpredictAPIRequest) SetParamOrderPredictQuery(_paramOrderPredictQuery *OrderPredictQuery) error {
	r._paramOrderPredictQuery = _paramOrderPredictQuery
	r.Set("param_order_predict_query", _paramOrderPredictQuery)
	return nil
}

// GetParamOrderPredictQuery ParamOrderPredictQuery Getter
func (r AlibabawdkscmlrporderpredictAPIRequest) GetParamOrderPredictQuery() *OrderPredictQuery {
	return r._paramOrderPredictQuery
}
