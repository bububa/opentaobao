package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
onecode代码通用鉴权 APIResponse
alibaba.interact.onecode.issue

手淘开放鉴权接口，仅用于tida接口鉴权，无输入输出。
*/
type AlibabaInteractOnecodeIssueAPIResponse struct {
    model.CommonResponse
    AlibabaInteractOnecodeIssueResponse
}

type AlibabaInteractOnecodeIssueResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_onecode_issue_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=0
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
