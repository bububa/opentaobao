package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
北京小蜜蜂配作业回传 API请求
alibaba.wdk.fulfill.dms.ebeecake.work.order.callback

北京小蜜蜂配作业回传。
*/
type AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest struct {
    model.Params
    // 作业单回传对象
    callbackOrder   *EbeecakeO2OCallbackOrder
}

// 初始化AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest对象
func NewAlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest() *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest{
    return &AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.dms.ebeecake.work.order.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest) SetCallbackOrder(callbackOrder *EbeecakeO2OCallbackOrder) error {
    r.callbackOrder = callbackOrder
    r.Set("callback_order", callbackOrder)
    return nil
}

// CallbackOrder Getter
func (r AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackRequest) GetCallbackOrder() *EbeecakeO2OCallbackOrder {
    return r.callbackOrder
}
