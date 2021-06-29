package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
作业小票查询接口 APIResponse
alibaba.wdk.fulfill.batch.query.by.batchids

根据节点等条件查询履约单小票信息
*/
type AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse struct {
    model.CommonResponse
    AlibabaWdkFulfillBatchQueryByBatchidsResponse
}

type AlibabaWdkFulfillBatchQueryByBatchidsResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_fulfill_batch_query_by_batchids_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果对象
    
    FulfillLogisticListResult   *FulfillLogisticListResult `json:"fulfill_logistic_list_result,omitempty" xml:"fulfill_logistic_list_result,omitempty"`

    
}
