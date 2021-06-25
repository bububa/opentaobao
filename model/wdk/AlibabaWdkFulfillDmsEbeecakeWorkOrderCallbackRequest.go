package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
北京小蜜蜂配作业回传 APIRequest
alibaba.wdk.fulfill.dms.ebeecake.work.order.callback

北京小蜜蜂配作业回传。
*/
type AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest struct {
    model.Params

    // 作业单回传对象
    callbackOrder   *EbeecakeO2OCallbackOrder 

}

func NewAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest() *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest{
    return &AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.dms.ebeecake.work.order.callback"
}

func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest) SetCallbackOrder(callbackOrder *EbeecakeO2OCallbackOrder) error {
    r.callbackOrder = callbackOrder
    r.Set("callback_order", callbackOrder)
    return nil
}

func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest) GetCallbackOrder() *EbeecakeO2OCallbackOrder {
    return r.callbackOrder
}

