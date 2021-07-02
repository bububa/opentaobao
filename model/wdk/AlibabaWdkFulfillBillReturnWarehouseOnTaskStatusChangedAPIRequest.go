package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest 退仓结果回传 API请求
// alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed
//
// 退货入仓结果回传
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest struct {
	model.Params
	// 退仓结果
	_returnWarehouseResult *ReturnWarehouseResult
}

// NewAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest 初始化AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest对象
func NewAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest() *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest {
	return &AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReturnWarehouseResult is ReturnWarehouseResult Setter
// 退仓结果
func (r *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest) SetReturnWarehouseResult(_returnWarehouseResult *ReturnWarehouseResult) error {
	r._returnWarehouseResult = _returnWarehouseResult
	r.Set("return_warehouse_result", _returnWarehouseResult)
	return nil
}

// GetReturnWarehouseResult ReturnWarehouseResult Getter
func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest) GetReturnWarehouseResult() *ReturnWarehouseResult {
	return r._returnWarehouseResult
}
