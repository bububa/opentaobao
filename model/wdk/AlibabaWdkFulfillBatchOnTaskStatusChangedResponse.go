package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流管控作业状态回传 APIResponse
alibaba.wdk.fulfill.batch.on.task.status.changed

物流管控作业状态回传
*/
type AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse struct {
    model.CommonResponse
    AlibabaWdkFulfillBatchOnTaskStatusChangedResponse
}

type AlibabaWdkFulfillBatchOnTaskStatusChangedResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_fulfill_batch_on_task_status_changed_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // fulfillLogisticSingleResult
    
    FulfillLogisticSingleResult   *FulfillLogisticDefaultResult `json:"fulfill_logistic_single_result,omitempty" xml:"fulfill_logistic_single_result,omitempty"`

    
}
