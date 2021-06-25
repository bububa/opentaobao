package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退仓结果回传 APIRequest
alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed

退货入仓结果回传
*/
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest struct {
    model.Params

    // 退仓结果
    returnWarehouseResult   *ReturnWarehouseResult 

}

func NewAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest() *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest{
    return &AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed"
}

func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest) SetReturnWarehouseResult(returnWarehouseResult *ReturnWarehouseResult) error {
    r.returnWarehouseResult = returnWarehouseResult
    r.Set("return_warehouse_result", returnWarehouseResult)
    return nil
}

func (r AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedRequest) GetReturnWarehouseResult() *ReturnWarehouseResult {
    return r.returnWarehouseResult
}

