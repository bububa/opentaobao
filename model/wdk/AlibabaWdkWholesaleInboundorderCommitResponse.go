package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
盒马帮退货信息回传接口 APIResponse
alibaba.wdk.wholesale.inboundorder.commit

盒马帮退货信息回传接口
*/
type AlibabaWdkWholesaleInboundorderCommitAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_wholesale_inboundorder_commit_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaWdkWholesaleInboundorderCommitApiResult `json:"result,omitempty" xml:"