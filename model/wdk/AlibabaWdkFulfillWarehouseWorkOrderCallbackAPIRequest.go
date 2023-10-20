package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest 标准化仓作业单回传接口 API请求
// alibaba.wdk.fulfill.warehouse.work.order.callback
//
// 标准化仓作业单回传接口
type AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *DrfHalfDayCcCallbackOrder
}

// NewAlibabaWdkFulfillWarehouseWorkOrderCallbackRequest 初始化AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest对象
func NewAlibabaWdkFulfillWarehouseWorkOrderCallbackRequest() *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest {
	return &AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) Reset() {
	r._callbackOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.warehouse.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *DrfHalfDayCcCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) GetCallbackOrder() *DrfHalfDayCcCallbackOrder {
	return r._callbackOrder
}

var poolAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillWarehouseWorkOrderCallbackRequest()
	},
}

// GetAlibabaWdkFulfillWarehouseWorkOrderCallbackRequest 从 sync.Pool 获取 AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest
func GetAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest() *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest {
	return poolAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest.Get().(*AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest)
}

// ReleaseAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest 将 AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest(v *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest.Put(v)
}
