package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest 顺丰仓作业回传 API请求
// alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback
//
// 顺丰仓作业单回传接口
type AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *SfB2CFmsCallbackOrder
}

// NewAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest 初始化AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest对象
func NewAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest() *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest {
	return &AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) Reset() {
	r._callbackOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *SfB2CFmsCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) GetCallbackOrder() *SfB2CFmsCallbackOrder {
	return r._callbackOrder
}

var poolAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest()
	},
}

// GetAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest 从 sync.Pool 获取 AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest
func GetAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest() *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest {
	return poolAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest.Get().(*AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest)
}

// ReleaseAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest 将 AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest(v *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest.Put(v)
}
