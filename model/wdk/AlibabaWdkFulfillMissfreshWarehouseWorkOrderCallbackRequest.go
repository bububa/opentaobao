package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
每日优鲜仓作业单回传接口 API请求
alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback

家乐福仓作业单回传接口
*/
type AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest struct {
    model.Params
    // 作业单回传对象
    _callbackOrder   *MissfreshO2OCallbackOrder
}

// 初始化AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest对象
func NewAlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest() *AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest{
    return &AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CallbackOrder Setter
// 作业单回传对象
func (r *AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest) SetCallbackOrder(_callbackOrder *MissfreshO2OCallbackOrder) error {
    r._callbackOrder = _callbackOrder
    r.Set("callback_order", _callbackOrder)
    return nil
}

// CallbackOrder Getter
func (r AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackRequest) GetCallbackOrder() *MissfreshO2OCallbackOrder {
    return r._callbackOrder
}
