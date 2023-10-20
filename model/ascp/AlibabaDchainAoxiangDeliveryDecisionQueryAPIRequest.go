package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangdeliverydecisionqueryAPIRequest 查询黑白名单快递 API请求
// alibaba.dchain.aoxiang.delivery.decision.query
//
// 查询黑白名单快递
type AlibabadchainaoxiangdeliverydecisionqueryAPIRequest struct {
	model.Params
	// 黑白名单快递查询入参
	_deliveryDecisionQueryRequest *DeliveryDecisionQueryRequest
}

// NewAlibabadchainaoxiangdeliverydecisionqueryRequest 初始化AlibabadchainaoxiangdeliverydecisionqueryAPIRequest对象
func NewAlibabadchainaoxiangdeliverydecisionqueryRequest() *AlibabadchainaoxiangdeliverydecisionqueryAPIRequest {
	return &AlibabadchainaoxiangdeliverydecisionqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangdeliverydecisionqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.delivery.decision.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangdeliverydecisionqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangdeliverydecisionqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryDecisionQueryRequest is DeliveryDecisionQueryRequest Setter
// 黑白名单快递查询入参
func (r *AlibabadchainaoxiangdeliverydecisionqueryAPIRequest) SetDeliveryDecisionQueryRequest(_deliveryDecisionQueryRequest *DeliveryDecisionQueryRequest) error {
	r._deliveryDecisionQueryRequest = _deliveryDecisionQueryRequest
	r.Set("delivery_decision_query_request", _deliveryDecisionQueryRequest)
	return nil
}

// GetDeliveryDecisionQueryRequest DeliveryDecisionQueryRequest Getter
func (r AlibabadchainaoxiangdeliverydecisionqueryAPIRequest) GetDeliveryDecisionQueryRequest() *DeliveryDecisionQueryRequest {
	return r._deliveryDecisionQueryRequest
}
