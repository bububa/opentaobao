package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
盒马帮退货信息回传接口 APIResponse
alibaba.wdk.wholesale.inboundorder.commit

盒马帮退货信息回传接口
*/
type AlibabaWdkWholesaleInboundorderCommitAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkWholesaleInboundorderCommitResponse `json:"alibaba_wdk_wholesale_inboundorder_commit_response,omitempty"`
}

type AlibabaWdkWholesaleInboundorderCommitResponse struct {

    // result
    Result   *AlibabaWdkWholesaleInboundorderCommitApiResult `json:"result,omitempty"`

}