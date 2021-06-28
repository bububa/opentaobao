package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销支付接口 APIResponse
alibaba.wdk.pos.trade.pay

轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易
*/
type AlibabaWdkPosTradePayAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkPosTradePayResponse `json:"alibaba_wdk_pos_trade_pay_response,omitempty"` 
    AlibabaWdkPosTradePayResponse
}

/* model for simplify = false
type AlibabaWdkPosTradePayResponse struct {

    // 支付结果
    
    Result  *struct {
        FastBuyPosPayResult  *FastBuyPosPayResult `json:"fast_buy_pos_pay_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkPosTradePayResponse struct {

    // 支付结果
    Result   *FastBuyPosPayResult `json:"result,omitempty"`

}
