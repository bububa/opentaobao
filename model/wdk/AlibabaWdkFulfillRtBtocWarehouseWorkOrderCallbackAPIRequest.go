package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大润发B2C仓作业单回传接口 API请求
alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback

大润发B2C仓作业单回传接口
*/
type AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest struct {
    model.Params
    // 作业单回传对象
    _callbackOrder   *DrfB2CCallbackOrder
}

// 初始化AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest对象
func NewAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackRequest() *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest{
    return &AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *DrfB2CCallbackOrder) error {
    r._callbackOrder = _callbackOrder
    r.Set("callback_order", _callbackOrder)
    return nil
}

// CallbackOrder Getter
func (r AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest) GetCallbackOrder() *DrfB2CCallbackOrder {
    return r._callbackOrder
}
