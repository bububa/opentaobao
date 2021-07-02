package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.dms.delivery.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *DeliveryCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// Get CallbackOrder Getter
func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest) GetCallbackOrder() *DeliveryCallbackOrder {
	return r._callbackOrder
}
