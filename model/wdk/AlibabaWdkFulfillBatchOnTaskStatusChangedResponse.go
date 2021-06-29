package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流管控作业状态回传 API返回值 
alibaba.wdk.fulfill.batch.on.task.status.changed

物流管控作业状态回传
*/
type AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse struct {
    model.CommonResponse
    AlibabaWdkFulfillBatchOnTaskStatusChangedResponse
}

// 物流管控作业状态回传 成功返回结果
type AlibabaWdkFulfillBatchOnTaskStatusChangedResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_fulfill_batch_on_task_status_changed_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // fulfillLogisticSingleResult
    FulfillLogisticSingleResult   *FulfillLogisticDefaultResult `json:"fulfill_logistic_single_result,omitempty" xml:"fulfill_logistic_single_result,omitempty"`
}
