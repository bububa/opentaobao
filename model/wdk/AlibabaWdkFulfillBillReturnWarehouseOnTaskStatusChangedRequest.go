package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退仓结果回传 API请求
alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed

退货入仓结果回传
*/
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest struct {
    model.Params
    // 退仓结果
    _returnWarehouseResult   *ReturnWarehouseResult
}

// 初始化AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest对象
func NewAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest() *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest{
    return &AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReturnWarehouseResult Setter
// 退仓结果
func (r *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest) SetReturnWarehouseResult(_returnWarehouseResult *ReturnWarehouseResult) error {
    r._returnWarehouseResult = _returnWarehouseResult
    r.Set("return_warehouse_result", _returnWarehouseResult)
    return nil
}

// ReturnWarehouseResult Getter
func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest) GetReturnWarehouseResult() *ReturnWarehouseResult {
    return r._returnWarehouseResult
}
