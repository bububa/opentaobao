package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest 中心仓查网格仓 API请求
// wdk.logistic.network.warehouse.resource.relation.query.from
//
// 盒马集市，中心仓查询网格仓
type WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest struct {
	model.Params
	// 查询参数
	_paramPageQueryWarehouseResourceRelationByFromRequest *PageQueryWarehouseResourceRelationByFromRequest
}

// NewWdkLogisticNetworkWarehouseResourceRelationQueryFromRequest 初始化WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest对象
func NewWdkLogisticNetworkWarehouseResourceRelationQueryFromRequest() *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest {
	return &WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) Reset() {
	r._paramPageQueryWarehouseResourceRelationByFromRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) GetApiMethodName() string {
	return "wdk.logistic.network.warehouse.resource.relation.query.from"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPageQueryWarehouseResourceRelationByFromRequest is ParamPageQueryWarehouseResourceRelationByFromRequest Setter
// 查询参数
func (r *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) SetParamPageQueryWarehouseResourceRelationByFromRequest(_paramPageQueryWarehouseResourceRelationByFromRequest *PageQueryWarehouseResourceRelationByFromRequest) error {
	r._paramPageQueryWarehouseResourceRelationByFromRequest = _paramPageQueryWarehouseResourceRelationByFromRequest
	r.Set("param_page_query_warehouse_resource_relation_by_from_request", _paramPageQueryWarehouseResourceRelationByFromRequest)
	return nil
}

// GetParamPageQueryWarehouseResourceRelationByFromRequest ParamPageQueryWarehouseResourceRelationByFromRequest Getter
func (r WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) GetParamPageQueryWarehouseResourceRelationByFromRequest() *PageQueryWarehouseResourceRelationByFromRequest {
	return r._paramPageQueryWarehouseResourceRelationByFromRequest
}

var poolWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkLogisticNetworkWarehouseResourceRelationQueryFromRequest()
	},
}

// GetWdkLogisticNetworkWarehouseResourceRelationQueryFromRequest 从 sync.Pool 获取 WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest
func GetWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest() *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest {
	return poolWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest.Get().(*WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest)
}

// ReleaseWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest 将 WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest 放入 sync.Pool
func ReleaseWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest(v *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest) {
	v.Reset()
	poolWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest.Put(v)
}
