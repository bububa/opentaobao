package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
是否为新支付订单 APIResponse
alibaba.mos.onsite.trade.isnewpayorder

退款时，老支付宝手淘退款接口需要查一下该订单是否为新支付订单
*/
type AlibabaMosOnsiteTradeIsnewpayorderAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMosOnsiteTradeIsnewpayorderResponse `json:"alibaba_mos_onsite_trade_isnewpayorder_response,omitempty"` 
    AlibabaMosOnsiteTradeIsnewpayorderResponse
}

/* model for simplify = false
type AlibabaMosOnsiteTradeIsnewpayorderResponse struct {

    // result
    
    Result  *struct {
        AlibabaMosOnsiteTradeIsnewpayorderResultDo  *AlibabaMosOnsiteTradeIsnewpayorderResultDo `json:"alibaba_mos_onsite_trade_isnewpayorder_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaMosOnsiteTradeIsnewpayorderResponse struct {

    // result
    Result   *AlibabaMosOnsiteTradeIsnewpayorderResultDo `json:"result,omitempty"`

}
