package traderate

import (
    "github.com/bububa/opentaobao/model"
)

/* 
针对父子订单新增批量评价 APIResponse
taobao.traderate.list.add

针对父子订单新增批量评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不用再通过该接口进行评价</font>)
*/
type TaobaoTraderateListAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTraderateListAddResponse `json:"traderate_list_add_response,omitempty"` 
    TaobaoTraderateListAddResponse
}

/* model for simplify = false
type TaobaoTraderateListAddResponse struct {

    // 返回的评论的信息，仅返回tid和created字段
    
    TradeRate  *struct {
        TradeRateRequest  *TradeRateRequest `json:"trade_rate_request,omitempty"`
    } `json:"trade_rate,omitempty"`
    

}
*/

type TaobaoTraderateListAddResponse struct {

    // 返回的评论的信息，仅返回tid和created字段
    TradeRate   *TradeRateRequest `json:"trade_rate,omitempty"`

}
