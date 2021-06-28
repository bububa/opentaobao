package traderate

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增单个评价 APIResponse
taobao.traderate.add

新增单个评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不能再通过该接口进行评价</font>)
*/
type TaobaoTraderateAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTraderateAddResponse `json:"traderate_add_response,omitempty"` 
    TaobaoTraderateAddResponse
}

/* model for simplify = false
type TaobaoTraderateAddResponse struct {

    // 返回tid、oid、create
    
    TradeRate  *struct {
        TradeRateRequest  *TradeRateRequest `json:"trade_rate_request,omitempty"`
    } `json:"trade_rate,omitempty"`
    

}
*/

type TaobaoTraderateAddResponse struct {

    // 返回tid、oid、create
    TradeRate   *TradeRateRequest `json:"trade_rate,omitempty"`

}
