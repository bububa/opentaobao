package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴权益投放接口 APIResponse
alibaba.latour.strategy.show

阿里巴巴权益平台权益投放接口
*/
type AlibabaLatourStrategyShowAPIResponse struct {
    model.CommonResponse
    Response *AlibabaLatourStrategyShowResponse `json:"alibaba_latour_strategy_show_response,omitempty"`
}

type AlibabaLatourStrategyShowResponse struct {

    // 返回结果
    Result   *AlibabaLatourStrategyShowResult `json:"result,omitempty"`

}
