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
    Response *AlibabaWdkWholesaleOutboundorderCommitResponse `json:"alibaba_wdk_wholesale_outboundorder_commit_response,omitempty"`
}

type AlibabaWdkWholesaleOutboundorderCommitResponse struct {

    // result
    Result   *AlibabaWdkWholesaleOutboundorderCommitApiResult `json:"result,omitempty"`

}
