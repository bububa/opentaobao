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
    AlibabaWdkWholesaleInboundorderCommitResponse
}

type AlibabaWdkWholesaleInboundorderCommitResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_wholesale_inboundorder_commit_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaWdkWholesaleInboundorderCommitApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
