package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest 大润发B2C仓作业单回传接口 API请求
// alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback
//
// 大润发B2C仓作业单回传接口
type AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *DrfB2CCallbackOrder
}

// NewAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest 初始化AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest对象
func NewAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest() *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest {
	return &AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) Reset() {
	r._callbackOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *DrfB2CCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) GetCallbackOrder() *DrfB2CCallbackOrder {
	return r._callbackOrder
}

var poolAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest()
	},
}

// GetAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest 从 sync.Pool 获取 AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest
func GetAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest() *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest {
	return poolAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest.Get().(*AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest)
}

// ReleaseAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest 将 AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest(v *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest.Put(v)
}
