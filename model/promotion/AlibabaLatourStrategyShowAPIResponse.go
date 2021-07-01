package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴权益投放接口 API返回值 
alibaba.latour.strategy.show

阿里巴巴权益平台权益投放接口
*/
type AlibabaLatourStrategyShowAPIResponse struct {
    model.CommonResponse
    AlibabaLatourStrategyShowAPIResponseModel
}

// 阿里巴巴权益投放接口 成功返回结果
type AlibabaLatourStrategyShowAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_latour_strategy_show_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *AlibabaLatourStrategyShowResult `json:"result,omitempty" xml:"result,omitempty"`
}
