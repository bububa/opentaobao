package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest 中心仓查网格仓 API请求
// wdk.logistic.network.warehouse.resource.relation.query.from
//
// 盒马集市，中心仓查询网格仓
type WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest struct {
	model.Params
	// 查询参数
	_paramPageQueryWarehouseResourceRelationByFromRequest *PageQueryWarehouseResourceRelationByFromRequest
}

// NewWdklogisticnetworkwarehouseresourcerelationqueryfromRequest 初始化WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest对象
func NewWdklogisticnetworkwarehouseresourcerelationqueryfromRequest() *WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest {
	return &WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest) GetApiMethodName() string {
	return "wdk.logistic.network.warehouse.resource.relation.query.from"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPageQueryWarehouseResourceRelationByFromRequest is ParamPageQueryWarehouseResourceRelationByFromRequest Setter
// 查询参数
func (r *WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest) SetParamPageQueryWarehouseResourceRelationByFromRequest(_paramPageQueryWarehouseResourceRelationByFromRequest *PageQueryWarehouseResourceRelationByFromRequest) error {
	r._paramPageQueryWarehouseResourceRelationByFromRequest = _paramPageQueryWarehouseResourceRelationByFromRequest
	r.Set("param_page_query_warehouse_resource_relation_by_from_request", _paramPageQueryWarehouseResourceRelationByFromRequest)
	return nil
}

// GetParamPageQueryWarehouseResourceRelationByFromRequest ParamPageQueryWarehouseResourceRelationByFromRequest Getter
func (r WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest) GetParamPageQueryWarehouseResourceRelationByFromRequest() *PageQueryWarehouseResourceRelationByFromRequest {
	return r._paramPageQueryWarehouseResourceRelationByFromRequest
}
