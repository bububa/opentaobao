package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
顺丰仓作业回传 APIRequest
alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback

顺丰仓作业单回传接口
*/
type AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest struct {
    model.Params

    // 作业单回传对象
    callbackOrder   *SfB2CFmsCallbackOrder 

}

func NewAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest() *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest{
    return &AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback"
}

func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest) SetCallbackOrder(callbackOrder *SfB2CFmsCallbackOrder) error {
    r.callbackOrder = callbackOrder
    r.Set("callback_order", callbackOrder)
    return nil
}

func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest) GetCallbackOrder() *SfB2CFmsCallbackOrder {
    return r.callbackOrder
}

