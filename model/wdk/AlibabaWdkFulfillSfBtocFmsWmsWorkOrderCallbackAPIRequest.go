package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *SfB2CFmsCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// Get CallbackOrder Getter
func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest) GetCallbackOrder() *SfB2CFmsCallbackOrder {
	return r._callbackOrder
}
