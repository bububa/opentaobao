package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkWarehouseOrderDispatchAPIRequest 仓作业下发 API请求
// wdk.warehouse.order.dispatch
//
// 牵牛花仓作业下发接口提供
type WdkWarehouseOrderDispatchAPIRequest struct {
	model.Params
	// 仓作业单
	_workOrder *WorkOrder
}

// NewWdkWarehouseOrderDispatchRequest 初始化WdkWarehouseOrderDispatchAPIRequest对象
func NewWdkWarehouseOrderDispatchRequest() *WdkWarehouseOrderDispatchAPIRequest {
	return &WdkWarehouseOrderDispatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkWarehouseOrderDispatchAPIRequest) GetApiMethodName() string {
	return "wdk.warehouse.order.dispatch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkWarehouseOrderDispatchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWorkOrder is WorkOrder Setter
// 仓作业单
func (r *WdkWarehouseOrderDispatchAPIRequest) SetWorkOrder(_workOrder *WorkOrder) error {
	r._workOrder = _workOrder
	r.Set("work_order", _workOrder)
	return nil
}

// GetWorkOrder WorkOrder Getter
func (r WdkWarehouseOrderDispatchAPIRequest) GetWorkOrder() *WorkOrder {
	return r._workOrder
}
