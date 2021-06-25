package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流管控作业状态回传 APIResponse
alibaba.wdk.fulfill.batch.on.task.status.changed

物流管控作业状态回传
*/
type AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkFulfillBatchOnTaskStatusChangedResponse `json:"alibaba_wdk_fulfill_batch_on_task_status_changed_response,omitempty"`
}

type AlibabaWdkFulfillBatchOnTaskStatusChangedResponse struct {

    // fulfillLogisticSingleResult
    FulfillLogisticSingleResult   *FulfillLogisticDefaultResult `json:"fulfill_logistic_single_result,omitempty"`

}
