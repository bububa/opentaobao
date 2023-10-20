package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest 每日优鲜仓作业单回传接口 API请求
// alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback
//
// 家乐福仓作业单回传接口
type AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *MissfreshO2OCallbackOrder
}

// NewAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest 初始化AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest对象
func NewAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest() *AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest {
	return &AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest) Reset() {
	r._callbackOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *MissfreshO2OCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest) GetCallbackOrder() *MissfreshO2OCallbackOrder {
	return r._callbackOrder
}

var poolAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest()
	},
}

// GetAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest 从 sync.Pool 获取 AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest
func GetAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest() *AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest {
	return poolAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest.Get().(*AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest)
}

// ReleaseAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest 将 AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest(v *AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest.Put(v)
}
