package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取前N有礼活动的开奖订单列表 APIResponse
taobao.trade.ordersort.get

获取前N有礼活动的开奖订单列表
*/
type TaobaoTradeOrdersortGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTradeOrdersortGetResponse `json:"trade_ordersort_get_response,omitempty"` 
    TaobaoTradeOrdersortGetResponse
}

/* model for simplify = false
type TaobaoTradeOrdersortGetResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoTradeOrdersortGetResult  *TaobaoTradeOrdersortGetResult `json:"taobao_trade_ordersort_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTradeOrdersortGetResponse struct {

    // 接口返回model
    Result   *TaobaoTradeOrdersortGetResult `json:"result,omitempty"`

}
