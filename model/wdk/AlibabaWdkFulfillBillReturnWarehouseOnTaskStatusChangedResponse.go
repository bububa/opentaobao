package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退仓结果回传 API返回值 
alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed

退货入仓结果回传
*/
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse struct {
    model.CommonResponse
    AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedResponse
}

// 退仓结果回传 成功返回结果
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_fulfill_bill_return_warehouse_on_task_status_changed_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // fulfillVoidResult
    Result   *FulfillVoidResult `json:"result,omitempty" xml:"result,omitempty"`
}
