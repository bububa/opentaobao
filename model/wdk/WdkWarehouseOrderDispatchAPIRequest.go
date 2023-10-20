package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkWarehouseOrderDispatchAPIRequest) Reset() {
	r._workOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkWarehouseOrderDispatchAPIRequest) GetApiMethodName() string {
	return "wdk.warehouse.order.dispatch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkWarehouseOrderDispatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkWarehouseOrderDispatchAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolWdkWarehouseOrderDispatchAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkWarehouseOrderDispatchRequest()
	},
}

// GetWdkWarehouseOrderDispatchRequest 从 sync.Pool 获取 WdkWarehouseOrderDispatchAPIRequest
func GetWdkWarehouseOrderDispatchAPIRequest() *WdkWarehouseOrderDispatchAPIRequest {
	return poolWdkWarehouseOrderDispatchAPIRequest.Get().(*WdkWarehouseOrderDispatchAPIRequest)
}

// ReleaseWdkWarehouseOrderDispatchAPIRequest 将 WdkWarehouseOrderDispatchAPIRequest 放入 sync.Pool
func ReleaseWdkWarehouseOrderDispatchAPIRequest(v *WdkWarehouseOrderDispatchAPIRequest) {
	v.Reset()
	poolWdkWarehouseOrderDispatchAPIRequest.Put(v)
}
