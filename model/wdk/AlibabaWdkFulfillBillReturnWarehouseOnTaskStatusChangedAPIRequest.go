package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest 退仓结果回传 API请求
// alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed
//
// 退货入仓结果回传
type AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest struct {
	model.Params
	// 退仓结果
	_returnWarehouseResult *ReturnWarehouseResult
}

// NewAlibabawdkfulfillbillreturnwarehouseontaskstatuschangedRequest 初始化AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest对象
func NewAlibabawdkfulfillbillreturnwarehouseontaskstatuschangedRequest() *AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest {
	return &AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReturnWarehouseResult is ReturnWarehouseResult Setter
// 退仓结果
func (r *AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest) SetReturnWarehouseResult(_returnWarehouseResult *ReturnWarehouseResult) error {
	r._returnWarehouseResult = _returnWarehouseResult
	r.Set("return_warehouse_result", _returnWarehouseResult)
	return nil
}

// GetReturnWarehouseResult ReturnWarehouseResult Getter
func (r AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest) GetReturnWarehouseResult() *ReturnWarehouseResult {
	return r._returnWarehouseResult
}
