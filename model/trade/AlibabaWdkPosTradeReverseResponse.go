package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销退款接口 APIResponse
alibaba.wdk.pos.trade.reverse

轻pos品牌营销场景，商家调用退款接口
*/
type AlibabaWdkPosTradeReverseAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkPosTradeReverseResponse `json:"alibaba_wdk_pos_trade_reverse_response,omitempty"`
}

type AlibabaWdkPosTradeReverseResponse struct {

    // 退款结果
    Result   *FastBuyPosReverseResult `json:"result,omitempty"`

}
