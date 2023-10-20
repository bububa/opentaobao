package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest 按网格仓查中心仓（带缓存） API请求
// wdk.logistic.network.warehouse.resource.relation.query.to.codes
//
// 盒马集市，网格仓查询中心仓
type WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest struct {
	model.Params
	// 入参
	_paramYxWarehouseResourceRelationQueryRequest *YxWarehouseResourceRelationQueryRequest
}

// NewWdklogisticnetworkwarehouseresourcerelationquerytocodesRequest 初始化WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest对象
func NewWdklogisticnetworkwarehouseresourcerelationquerytocodesRequest() *WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest {
	return &WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest) GetApiMethodName() string {
	return "wdk.logistic.network.warehouse.resource.relation.query.to.codes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamYxWarehouseResourceRelationQueryRequest is ParamYxWarehouseResourceRelationQueryRequest Setter
// 入参
func (r *WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest) SetParamYxWarehouseResourceRelationQueryRequest(_paramYxWarehouseResourceRelationQueryRequest *YxWarehouseResourceRelationQueryRequest) error {
	r._paramYxWarehouseResourceRelationQueryRequest = _paramYxWarehouseResourceRelationQueryRequest
	r.Set("param_yx_warehouse_resource_relation_query_request", _paramYxWarehouseResourceRelationQueryRequest)
	return nil
}

// GetParamYxWarehouseResourceRelationQueryRequest ParamYxWarehouseResourceRelationQueryRequest Getter
func (r WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest) GetParamYxWarehouseResourceRelationQueryRequest() *YxWarehouseResourceRelationQueryRequest {
	return r._paramYxWarehouseResourceRelationQueryRequest
}
