package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标准化仓作业单回传接口 API请求
alibaba.wdk.fulfill.warehouse.work.order.callback

标准化仓作业单回传接口
*/
type AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest struct {
    model.Params
    // 作业单回传对象
    _callbackOrder   *DrfHalfDayCcCallbackOrder
}

// 初始化AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest对象
func NewAlibabaWdkFulfillWarehouseWorkOrderCallbackRequest() *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest{
    return &AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.warehouse.work.order.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) SetCallbackOrder(_callbackOrder *DrfHalfDayCcCallbackOrder) error {
    r._callbackOrder = _callbackOrder
    r.Set("callback_order", _callbackOrder)
    return nil
}

// CallbackOrder Getter
func (r AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest) GetCallbackOrder() *DrfHalfDayCcCallbackOrder {
    return r._callbackOrder
}
