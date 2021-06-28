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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_onecode_issue_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=0
    
    Result   string `json:"result,omitempty" xml:"