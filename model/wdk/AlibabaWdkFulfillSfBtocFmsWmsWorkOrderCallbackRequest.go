package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
顺丰仓作业回传 API请求
alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback

顺丰仓作业单回传接口
*/
type AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest struct {
    model.Params
    // 作业单回传对象
    callbackOrder   *SfB2CFmsCallbackOrder
}

// 初始化AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest对象
func NewAlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest() *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest{
    return &AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest) SetCallbackOrder(callbackOrder *SfB2CFmsCallbackOrder) error {
    r.callbackOrder = callbackOrder
    r.Set("callback_order", callbackOrder)
    return nil
}

// CallbackOrder Getter
func (r AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackRequest) GetCallbackOrder() *SfB2CFmsCallbackOrder {
    return r.callbackOrder
}
