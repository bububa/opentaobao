package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
每日优鲜仓作业单回传接口 APIRequest
alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback

家乐福仓作业单回传接口
*/
type AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest struct {
    model.Params

    // 作业单回传对象
    callbackOrder   *MissfreshO2OCallbackOrder 

}

func NewAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest() *AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest{
    return &AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback"
}

func (r AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest) SetCallbackOrder(callbackOrder *MissfreshO2OCallbackOrder) error {
    r.callbackOrder = callbackOrder
    r.Set("callback_order", callbackOrder)
    return nil
}

func (r AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest) GetCallbackOrder() *MissfreshO2OCallbackOrder {
    return r.callbackOrder
}

