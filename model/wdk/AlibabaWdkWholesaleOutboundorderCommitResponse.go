package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
盒马帮发货信息回传接口 APIResponse
alibaba.wdk.wholesale.outboundorder.commit

盒马帮发货信息回传接口
*/
type AlibabaWdkWholesaleOutboundorderCommitAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkWholesaleOutboundorderCommitResponse `json:"alibaba_wdk_wholesale_outboundorder_commit_response,omitempty"` 
    AlibabaWdkWholesaleOutboundorderCommitResponse
}

/* model for simplify = false
type AlibabaWdkWholesaleOutboundorderCommitResponse struct {

    // result
    
    Result  *struct {
        AlibabaWdkWholesaleOutboundorderCommitApiResult  *AlibabaWdkWholesaleOutboundorderCommitApiResult `json:"alibaba_wdk_wholesale_outboundorder_commit_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkWholesaleOutboundorderCommitResponse struct {

    // result
    Result   *AlibabaWdkWholesaleOutboundorderCommitApiResult `json:"result,omitempty"`

}
