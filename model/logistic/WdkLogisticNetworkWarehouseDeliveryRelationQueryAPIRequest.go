package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest 仓站（网格仓自提点）关系查询 API请求
// wdk.logistic.network.warehouse.delivery.relation.query
//
// 盒马集市，仓站（网格仓自提点）关系查询
type WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest struct {
	model.Params
	// 参数
	_paramWarehouseDeliveryRelationPageQueryRequest *WarehouseDeliveryRelationPageQueryRequest
}

// NewWdkLogisticNetworkWarehouseDeliveryRelationQueryRequest 初始化WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest对象
func NewWdkLogisticNetworkWarehouseDeliveryRelationQueryRequest() *WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest {
	return &WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) GetApiMethodName() string {
	return "wdk.logistic.network.warehouse.delivery.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamWarehouseDeliveryRelationPageQueryRequest Setter
// 参数
func (r *WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) SetParamWarehouseDeliveryRelationPageQueryRequest(_paramWarehouseDeliveryRelationPageQueryRequest *WarehouseDeliveryRelationPageQueryRequest) error {
	r._paramWarehouseDeliveryRelationPageQueryRequest = _paramWarehouseDeliveryRelationPageQueryRequest
	r.Set("param_warehouse_delivery_relation_page_query_request", _paramWarehouseDeliveryRelationPageQueryRequest)
	return nil
}

// Get ParamWarehouseDeliveryRelationPageQueryRequest Getter
func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) GetParamWarehouseDeliveryRelationPageQueryRequest() *WarehouseDeliveryRelationPageQueryRequest {
	return r._paramWarehouseDeliveryRelationPageQueryRequest
}
