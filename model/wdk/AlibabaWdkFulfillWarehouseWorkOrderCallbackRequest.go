package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标准化仓作业单回传接口 APIRequest
alibaba.wdk.fulfill.warehouse.work.order.callback

标准化仓作业单回传接口
*/
type AlibabaWdkFulfillWarehouseWorkOrderCallbackRequest struct {
    model.Params

    // 作业单回传对象
    callbackOrder   *DrfHalfDayCcCallbackOrder 

}

func NewAlibabaWdkFulfillWarehouseWorkOrderCallbackRequest() *AlibabaWdkFulfillWarehouseWorkOrderCallbackRequest{
    return &AlibabaWdkFulfillWarehouseWorkOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillWarehouseWorkOrderCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.warehouse.work.order.callback"
}

func (r AlibabaWdkFulfillWarehouseWorkOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillWarehouseWorkOrderCallbackRequest) SetCallbackOrder(callbackOrder *DrfHalfDayCcCallbackOrder) error {
    r.callbackOrder = callbackOrder
    r.Set("callback_order", callbackOrder)
    return nil
}

func (r AlibabaWdkFulfillWarehouseWorkOrderCallbackRequest) GetCallbackOrder() *DrfHalfDayCcCallbackOrder {
    return r.callbackOrder
}

