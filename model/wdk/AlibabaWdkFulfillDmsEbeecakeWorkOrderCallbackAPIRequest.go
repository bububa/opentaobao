package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.dms.ebeecake.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *EbeecakeO2OCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// Get CallbackOrder Getter
func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest) GetCallbackOrder() *EbeecakeO2OCallbackOrder {
	return r._callbackOrder
}
