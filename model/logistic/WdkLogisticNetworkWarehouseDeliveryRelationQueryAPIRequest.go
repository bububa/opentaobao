package logistic

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) Reset() {
	r._paramWarehouseDeliveryRelationPageQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) GetApiMethodName() string {
	return "wdk.logistic.network.warehouse.delivery.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWarehouseDeliveryRelationPageQueryRequest is ParamWarehouseDeliveryRelationPageQueryRequest Setter
// 参数
func (r *WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) SetParamWarehouseDeliveryRelationPageQueryRequest(_paramWarehouseDeliveryRelationPageQueryRequest *WarehouseDeliveryRelationPageQueryRequest) error {
	r._paramWarehouseDeliveryRelationPageQueryRequest = _paramWarehouseDeliveryRelationPageQueryRequest
	r.Set("param_warehouse_delivery_relation_page_query_request", _paramWarehouseDeliveryRelationPageQueryRequest)
	return nil
}

// GetParamWarehouseDeliveryRelationPageQueryRequest ParamWarehouseDeliveryRelationPageQueryRequest Getter
func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) GetParamWarehouseDeliveryRelationPageQueryRequest() *WarehouseDeliveryRelationPageQueryRequest {
	return r._paramWarehouseDeliveryRelationPageQueryRequest
}

var poolWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkLogisticNetworkWarehouseDeliveryRelationQueryRequest()
	},
}

// GetWdkLogisticNetworkWarehouseDeliveryRelationQueryRequest 从 sync.Pool 获取 WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest
func GetWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest() *WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest {
	return poolWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest.Get().(*WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest)
}

// ReleaseWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest 将 WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest 放入 sync.Pool
func ReleaseWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest(v *WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest) {
	v.Reset()
	poolWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest.Put(v)
}
