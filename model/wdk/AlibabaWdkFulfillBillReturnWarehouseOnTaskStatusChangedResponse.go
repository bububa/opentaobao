package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退仓结果回传 APIResponse
alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed

退货入仓结果回传
*/
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedResponse `json:"alibaba_wdk_fulfill_bill_return_warehouse_on_task_status_changed_response,omitempty"` 
    AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedResponse
}

/* model for simplify = false
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedResponse struct {

    // fulfillVoidResult
    
    Result  *struct {
        FulfillVoidResult  *FulfillVoidResult `json:"fulfill_void_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedResponse struct {

    // fulfillVoidResult
    Result   *FulfillVoidResult `json:"result,omitempty"`

}
