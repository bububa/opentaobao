package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴权益发放接口 API返回值 
alibaba.latour.strategy.issue

阿里巴巴权益平台权益发放接口
*/
type AlibabaLatourStrategyIssueAPIResponse struct {
    model.CommonResponse
    AlibabaLatourStrategyIssueResponse
}

// 阿里巴巴权益发放接口 成功返回结果
type AlibabaLatourStrategyIssueResponse struct {
    XMLName xml.Name `xml:"alibaba_latour_strategy_issue_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *AlibabaLatourStrategyIssueResult `json:"result,omitempty" xml:"result,omitempty"`
}
