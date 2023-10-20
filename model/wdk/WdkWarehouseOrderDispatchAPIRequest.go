package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkwarehouseorderdispatchAPIRequest 仓作业下发 API请求
// wdk.warehouse.order.dispatch
//
// 牵牛花仓作业下发接口提供
type WdkwarehouseorderdispatchAPIRequest struct {
	model.Params
	// 仓作业单
	_workOrder *WorkOrder
}

// NewWdkwarehouseorderdispatchRequest 初始化WdkwarehouseorderdispatchAPIRequest对象
func NewWdkwarehouseorderdispatchRequest() *WdkwarehouseorderdispatchAPIRequest {
	return &WdkwarehouseorderdispatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkwarehouseorderdispatchAPIRequest) GetApiMethodName() string {
	return "wdk.warehouse.order.dispatch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkwarehouseorderdispatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkwarehouseorderdispatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkOrder is WorkOrder Setter
// 仓作业单
func (r *WdkwarehouseorderdispatchAPIRequest) SetWorkOrder(_workOrder *WorkOrder) error {
	r._workOrder = _workOrder
	r.Set("work_order", _workOrder)
	return nil
}

// GetWorkOrder WorkOrder Getter
func (r WdkwarehouseorderdispatchAPIRequest) GetWorkOrder() *WorkOrder {
	return r._workOrder
}
