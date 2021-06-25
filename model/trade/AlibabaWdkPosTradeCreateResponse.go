package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销下单接口 APIResponse
alibaba.wdk.pos.trade.create

提供给石基进行轻pos品牌营销下单
*/
type AlibabaWdkPosTradeCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkPosTradeCreateResponse `json:"alibaba_wdk_pos_trade_create_response,omitempty"`
}

type AlibabaWdkPosTradeCreateResponse struct {

    // 创单结果
    Result   *FastBuyPosCreateResult `json:"result,omitempty"`

}
