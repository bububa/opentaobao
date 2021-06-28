package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改推广单元 APIResponse
alibaba.scbp.ad.group.update.ad.group.batch

修改推广单元
*/
type AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdGroupUpdateAdGroupBatchResponse `json:"alibaba_scbp_ad_group_update_ad_group_batch_response,omitempty"` 
    AlibabaScbpAdGroupUpdateAdGroupBatchResponse
}

/* model for simplify = false
type AlibabaScbpAdGroupUpdateAdGroupBatchResponse struct {

    // 返回结果
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdGroupUpdateAdGroupBatchResponse struct {

    // 返回结果
    Result   int64 `json:"result,omitempty"`

}
