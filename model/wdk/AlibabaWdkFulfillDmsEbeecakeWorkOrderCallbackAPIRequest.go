package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest 北京小蜜蜂配作业回传 API请求
// alibaba.wdk.fulfill.dms.ebeecake.work.order.callback
//
// 北京小蜜蜂配作业回传。
type AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *EbeecakeO2OCallbackOrder
}

// NewAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest 初始化AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest对象
func NewAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest() *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest {
	return &AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) Reset() {
	r._callbackOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.dms.ebeecake.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *EbeecakeO2OCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) GetCallbackOrder() *EbeecakeO2OCallbackOrder {
	return r._callbackOrder
}

var poolAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest()
	},
}

// GetAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest 从 sync.Pool 获取 AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest
func GetAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest() *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest {
	return poolAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest.Get().(*AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest)
}

// ReleaseAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest 将 AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest(v *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest.Put(v)
}
