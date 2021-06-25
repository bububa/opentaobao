package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
作业小票查询接口 APIResponse
alibaba.wdk.fulfill.batch.query.by.batchids

根据节点等条件查询履约单小票信息
*/
type AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkFulfillBatchQueryByBatchidsResponse `json:"alibaba_wdk_fulfill_batch_query_by_batchids_response,omitempty"`
}

type AlibabaWdkFulfillBatchQueryByBatchidsResponse struct {

    // 查询结果对象
    FulfillLogisticListResult   *FulfillLogisticListResult `json:"fulfill_logistic_list_result,omitempty"`

}
