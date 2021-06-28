package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下新退款接口（专为老退款接口调用） APIResponse
alibaba.mos.onsite.trade.oldrefund

线下新退款接口（专为老退款接口调用）。新接口支付支付宝、手淘、天猫，老接口退款时，需要调用该接口退新单，并适配老接口响应参数返回
*/
type AlibabaMosOnsiteTradeOldrefundAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMosOnsiteTradeOldrefundResponse `json:"alibaba_mos_onsite_trade_oldrefund_response,omitempty"` 
    AlibabaMosOnsiteTradeOldrefundResponse
}

/* model for simplify = false
type AlibabaMosOnsiteTradeOldrefundResponse struct {

    // 交易退款响应
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaMosOnsiteTradeOldrefundResponse struct {

    // 交易退款响应
    Result   string `json:"result,omitempty"`

}
