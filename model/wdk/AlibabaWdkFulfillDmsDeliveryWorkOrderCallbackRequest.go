package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
末端配配送作业回传 APIRequest
alibaba.wdk.fulfill.dms.delivery.work.order.callback

末端配配送作业回传。
*/
type AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest struct {
    model.Params

    // 作业单回传对象
    callbackOrder   *DeliveryCallbackOrder 

}

func NewAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest() *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest{
    return &AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.dms.delivery.work.order.callback"
}

func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest) SetCallbackOrder(callbackOrder *DeliveryCallbackOrder) error {
    r.callbackOrder = callbackOrder
    r.Set("callback_order", callbackOrder)
    return nil
}

func (r AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackRequest) GetCallbackOrder() *DeliveryCallbackOrder {
    return r.callbackOrder
}

