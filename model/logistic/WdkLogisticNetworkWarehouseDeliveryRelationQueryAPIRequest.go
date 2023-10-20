package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest 仓站（网格仓自提点）关系查询 API请求
// wdk.logistic.network.warehouse.delivery.relation.query
//
// 盒马集市，仓站（网格仓自提点）关系查询
type WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest struct {
	model.Params
	// 参数
	_paramWarehouseDeliveryRelationPageQueryRequest *WarehouseDeliveryRelationPageQueryRequest
}

// NewWdklogisticnetworkwarehousedeliveryrelationqueryRequest 初始化WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest对象
func NewWdklogisticnetworkwarehousedeliveryrelationqueryRequest() *WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest {
	return &WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest) GetApiMethodName() string {
	return "wdk.logistic.network.warehouse.delivery.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWarehouseDeliveryRelationPageQueryRequest is ParamWarehouseDeliveryRelationPageQueryRequest Setter
// 参数
func (r *WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest) SetParamWarehouseDeliveryRelationPageQueryRequest(_paramWarehouseDeliveryRelationPageQueryRequest *WarehouseDeliveryRelationPageQueryRequest) error {
	r._paramWarehouseDeliveryRelationPageQueryRequest = _paramWarehouseDeliveryRelationPageQueryRequest
	r.Set("param_warehouse_delivery_relation_page_query_request", _paramWarehouseDeliveryRelationPageQueryRequest)
	return nil
}

// GetParamWarehouseDeliveryRelationPageQueryRequest ParamWarehouseDeliveryRelationPageQueryRequest Getter
func (r WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest) GetParamWarehouseDeliveryRelationPageQueryRequest() *WarehouseDeliveryRelationPageQueryRequest {
	return r._paramWarehouseDeliveryRelationPageQueryRequest
}
