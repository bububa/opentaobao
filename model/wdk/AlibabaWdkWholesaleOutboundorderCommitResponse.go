package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
盒马帮发货信息回传接口 APIResponse
alibaba.wdk.wholesale.outboundorder.commit

盒马帮发货信息回传接口
*/
type AlibabaWdkWholesaleOutboundorderCommitAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_wholesale_outboundorder_commit_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaWdkWholesaleOutboundorderCommitApiResult `json:"result,omitempty" xml:"