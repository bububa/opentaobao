package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新商场当面付交易查询 APIResponse
alibaba.mos.onsite.trade.query

本接口提供新商场当面付订单的查询的功能，商户可以通过本接口主动查询订单状态，完成下一步的业务逻辑。
商户系统应在两种场景下调用此接口: 商户POS系统应该在调用[条码支付请求接口]并成功返回后，调用此接口查询订单的处理状态。
*/
type AlibabaMosOnsiteTradeQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosOnsiteTradeQueryResponse `json:"alibaba_mos_onsite_trade_query_response,omitempty"`
}

type AlibabaMosOnsiteTradeQueryResponse struct {

    // 查询结果对象。必然返回
    OnsiteTradeQueryResponse   *OnsiteTradeQueryResponse `json:"onsite_trade_query_response,omitempty"`

}
