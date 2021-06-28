package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建推广单元 APIResponse
alibaba.scbp.ad.group.create.ad.group.batch

创建推广单元
*/
type AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdGroupCreateAdGroupBatchResponse `json:"alibaba_scbp_ad_group_create_ad_group_batch_response,omitempty"` 
    AlibabaScbpAdGroupCreateAdGroupBatchResponse
}

/* model for simplify = false
type AlibabaScbpAdGroupCreateAdGroupBatchResponse struct {

    // 返回结果
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdGroupCreateAdGroupBatchResponse struct {

    // 返回结果
    Result   string `json:"result,omitempty"`

}
