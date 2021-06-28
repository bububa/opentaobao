package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退款查询 APIResponse
alibaba.mos.onsite.trade.queryrefund

商户可使用该接口查询退款请求是否执行成功。
*/
type AlibabaMosOnsiteTradeQueryrefundAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMosOnsiteTradeQueryrefundResponse `json:"alibaba_mos_onsite_trade_queryrefund_response,omitempty"` 
    AlibabaMosOnsiteTradeQueryrefundResponse
}

/* model for simplify = false
type AlibabaMosOnsiteTradeQueryrefundResponse struct {

    // result
    
    Result  *struct {
        AlibabaMosOnsiteTradeQueryrefundResultDo  *AlibabaMosOnsiteTradeQueryrefundResultDo `json:"alibaba_mos_onsite_trade_queryrefund_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaMosOnsiteTradeQueryrefundResponse struct {

    // result
    Result   *AlibabaMosOnsiteTradeQueryrefundResultDo `json:"result,omitempty"`

}
