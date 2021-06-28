package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴权益发放接口 APIResponse
alibaba.latour.strategy.issue

阿里巴巴权益平台权益发放接口
*/
type AlibabaLatourStrategyIssueAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_latour_strategy_issue_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaLatourStrategyIssueResult `json:"result,omitempty" xml:"