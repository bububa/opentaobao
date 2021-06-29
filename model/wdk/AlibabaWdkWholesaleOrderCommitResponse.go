package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
盒马帮采购确认订单接口 APIResponse
alibaba.wdk.wholesale.order.commit

盒马帮采购确认订单接口
*/
type AlibabaWdkWholesaleOrderCommitAPIResponse struct {
    model.CommonResponse
    AlibabaWdkWholesaleOrderCommitResponse
}

type AlibabaWdkWholesaleOrderCommitResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_wholesale_order_commit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkWholesaleOrderCommitApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
