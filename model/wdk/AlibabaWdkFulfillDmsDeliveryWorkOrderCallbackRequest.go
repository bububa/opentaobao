package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
末端配配送作业回传 API请求
alibaba.wdk.fulfill.dms.delivery.work.order.callback

末端配配送作业回传。
*/
type AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest struct {
    model.Params
    // 作业单回传对象
    callbackOrder   *DeliveryCallbackOrder
}

// 初始化AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest对象
func NewAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest() *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest{
    return &AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.dms.delivery.work.order.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest) SetCallbackOrder(callbackOrder *DeliveryCallbackOrder) error {
    r.callbackOrder = callbackOrder
    r.Set("callback_order", callbackOrder)
    return nil
}

// CallbackOrder Getter
func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest) GetCallbackOrder() *DeliveryCallbackOrder {
    return r.callbackOrder
}
