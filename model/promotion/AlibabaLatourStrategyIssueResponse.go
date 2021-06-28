package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴权益发放接口 APIResponse
alibaba.latour.strategy.issue

阿里巴巴权益平台权益发放接口
*/
type AlibabaLatourStrategyIssueAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLatourStrategyIssueResponse `json:"alibaba_latour_strategy_issue_response,omitempty"` 
    AlibabaLatourStrategyIssueResponse
}

/* model for simplify = false
type AlibabaLatourStrategyIssueResponse struct {

    // 返回结果
    
    Result  *struct {
        AlibabaLatourStrategyIssueResult  *AlibabaLatourStrategyIssueResult `json:"alibaba_latour_strategy_issue_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaLatourStrategyIssueResponse struct {

    // 返回结果
    Result   *AlibabaLatourStrategyIssueResult `json:"result,omitempty"`

}
