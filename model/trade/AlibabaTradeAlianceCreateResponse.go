package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推客平台订单回流 APIResponse
alibaba.trade.aliance.create

推客平台订单回流
*/
type AlibabaTradeAlianceCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTradeAlianceCreateResponse `json:"alibaba_trade_aliance_create_response,omitempty"` 
    AlibabaTradeAlianceCreateResponse
}

/* model for simplify = false
type AlibabaTradeAlianceCreateResponse struct {

    // 订单创建结果
    
    Result  *struct {
        AlibabaTradeAlianceCreateResultModel  *AlibabaTradeAlianceCreateResultModel `json:"alibaba_trade_aliance_create_result_model,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaTradeAlianceCreateResponse struct {

    // 订单创建结果
    Result   *AlibabaTradeAlianceCreateResultModel `json:"result,omitempty"`

}
