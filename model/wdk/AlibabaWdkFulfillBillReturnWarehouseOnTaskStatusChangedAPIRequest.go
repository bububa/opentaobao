package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest) Reset() {
	r._returnWarehouseResult = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest()
	},
}

// GetAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest 从 sync.Pool 获取 AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest
func GetAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest() *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest {
	return poolAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest.Get().(*AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest)
}

// ReleaseAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest 将 AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest(v *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest.Put(v)
}
