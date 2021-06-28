package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
onecode代码通用鉴权 APIResponse
alibaba.interact.onecode.issue

手淘开放鉴权接口，仅用于tida接口鉴权，无输入输出。
*/
type AlibabaInteractOnecodeIssueAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractOnecodeIssueResponse `json:"alibaba_interact_onecode_issue_response,omitempty"` 
    AlibabaInteractOnecodeIssueResponse
}

/* model for simplify = false
type AlibabaInteractOnecodeIssueResponse struct {

    // result=0
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractOnecodeIssueResponse struct {

    // result=0
    Result   string `json:"result,omitempty"`

}
