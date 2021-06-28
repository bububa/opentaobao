package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除推广单元 APIResponse
alibaba.scbp.ad.group.delete.ad.group.batch

删除推广单元
*/
type AlibabaScbpAdGroupDeleteAdGroupBatchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdGroupDeleteAdGroupBatchResponse `json:"alibaba_scbp_ad_group_delete_ad_group_batch_response,omitempty"` 
    AlibabaScbpAdGroupDeleteAdGroupBatchResponse
}

/* model for simplify = false
type AlibabaScbpAdGroupDeleteAdGroupBatchResponse struct {

    // 返回结果
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdGroupDeleteAdGroupBatchResponse struct {

    // 返回结果
    Result   int64 `json:"result,omitempty"`

}
