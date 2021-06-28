package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退仓结果回传 APIResponse
alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed

退货入仓结果回传
*/
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse struct {
    model.CommonResponse
    AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedResponse
}

type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_fulfill_bill_return_warehouse_on_task_status_changed_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // fulfillVoidResult
    
    Result   *FulfillVoidResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
