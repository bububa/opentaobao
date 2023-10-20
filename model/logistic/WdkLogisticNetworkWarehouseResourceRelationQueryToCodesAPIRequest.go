package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest 按网格仓查中心仓（带缓存） API请求
// wdk.logistic.network.warehouse.resource.relation.query.to.codes
//
// 盒马集市，网格仓查询中心仓
type WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest struct {
	model.Params
	// 入参
	_paramYxWarehouseResourceRelationQueryRequest *YxWarehouseResourceRelationQueryRequest
}

// NewWdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest 初始化WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest对象
func NewWdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest() *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest {
	return &WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) Reset() {
	r._paramYxWarehouseResourceRelationQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) GetApiMethodName() string {
	return "wdk.logistic.network.warehouse.resource.relation.query.to.codes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamYxWarehouseResourceRelationQueryRequest is ParamYxWarehouseResourceRelationQueryRequest Setter
// 入参
func (r *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) SetParamYxWarehouseResourceRelationQueryRequest(_paramYxWarehouseResourceRelationQueryRequest *YxWarehouseResourceRelationQueryRequest) error {
	r._paramYxWarehouseResourceRelationQueryRequest = _paramYxWarehouseResourceRelationQueryRequest
	r.Set("param_yx_warehouse_resource_relation_query_request", _paramYxWarehouseResourceRelationQueryRequest)
	return nil
}

// GetParamYxWarehouseResourceRelationQueryRequest ParamYxWarehouseResourceRelationQueryRequest Getter
func (r WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) GetParamYxWarehouseResourceRelationQueryRequest() *YxWarehouseResourceRelationQueryRequest {
	return r._paramYxWarehouseResourceRelationQueryRequest
}

var poolWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest()
	},
}

// GetWdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest 从 sync.Pool 获取 WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest
func GetWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest() *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest {
	return poolWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest.Get().(*WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest)
}

// ReleaseWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest 将 WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest 放入 sync.Pool
func ReleaseWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest(v *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) {
	v.Reset()
	poolWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest.Put(v)
}
