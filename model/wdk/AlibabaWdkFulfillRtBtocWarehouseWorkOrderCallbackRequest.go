package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大润发B2C仓作业单回传接口 APIRequest
alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback

大润发B2C仓作业单回传接口
*/
type AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest struct {
    model.Params

    // 作业单回传对象
    callbackOrder   *DrfB2CCallbackOrder 

}

func NewAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest() *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest{
    return &AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback"
}

func (r AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest) SetCallbackOrder(callbackOrder *DrfB2CCallbackOrder) error {
    r.callbackOrder = callbackOrder
    r.Set("callback_order", callbackOrder)
    return nil
}

func (r AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest) GetCallbackOrder() *DrfB2CCallbackOrder {
    return r.callbackOrder
}

