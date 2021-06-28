package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新商场当面付商户扫码付 APIResponse
alibaba.mos.onsite.trade.pay

收银员使用扫码设备读取用户“付款码”后，将二维码或条码信息通过本接口上送至喵街发起支付。
*/
type AlibabaMosOnsiteTradePayAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMosOnsiteTradePayResponse `json:"alibaba_mos_onsite_trade_pay_response,omitempty"` 
    AlibabaMosOnsiteTradePayResponse
}

/* model for simplify = false
type AlibabaMosOnsiteTradePayResponse struct {

    // 创建订单响应信息。必然返回
    
    OnsiteTradePayResponse  *struct {
        OnsiteTradePayResponse  *OnsiteTradePayResponse `json:"onsite_trade_pay_response,omitempty"`
    } `json:"onsite_trade_pay_response,omitempty"`
    

}
*/

type AlibabaMosOnsiteTradePayResponse struct {

    // 创建订单响应信息。必然返回
    OnsiteTradePayResponse   *OnsiteTradePayResponse `json:"onsite_trade_pay_response,omitempty"`

}
