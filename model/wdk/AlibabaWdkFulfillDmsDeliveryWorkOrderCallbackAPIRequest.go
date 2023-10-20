package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest 末端配配送作业回传 API请求
// alibaba.wdk.fulfill.dms.delivery.work.order.callback
//
// 末端配配送作业回传。
type AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *DeliveryCallbackOrder
}

// NewAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest 初始化AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest对象
func NewAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest() *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest {
	return &AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) Reset() {
	r._callbackOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.dms.delivery.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *DeliveryCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) GetCallbackOrder() *DeliveryCallbackOrder {
	return r._callbackOrder
}

var poolAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest()
	},
}

// GetAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest 从 sync.Pool 获取 AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest
func GetAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest() *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest {
	return poolAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest.Get().(*AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest)
}

// ReleaseAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest 将 AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest(v *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest.Put(v)
}
