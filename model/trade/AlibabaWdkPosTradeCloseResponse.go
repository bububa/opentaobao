package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销关单接口 APIResponse
alibaba.wdk.pos.trade.close

轻pos品牌营销场景，提供关单接口给外部商家
*/
type AlibabaWdkPosTradeCloseAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkPosTradeCloseResponse `json:"alibaba_wdk_pos_trade_close_response,omitempty"`
}

type AlibabaWdkPosTradeCloseResponse struct {

    // 关单结果
    Result   *FastBuyPosCloseResult `json:"result,omitempty"`

}
