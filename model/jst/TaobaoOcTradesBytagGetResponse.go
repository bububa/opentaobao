package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
标签查询订单 APIResponse
taobao.oc.trades.bytag.get

根据标签查询订单编号
*/
type TaobaoOcTradesBytagGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOcTradesBytagGetResponse `json:"oc_trades_bytag_get_response,omitempty"` 
    TaobaoOcTradesBytagGetResponse
}

/* model for simplify = false
type TaobaoOcTradesBytagGetResponse struct {

    // 打了该标签的订单编号列表
    
    Tids  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"tids,omitempty"`
    

    // 总数
    
    Totals   int64 `json:"totals,omitempty"`
    

}
*/

type TaobaoOcTradesBytagGetResponse struct {

    // 打了该标签的订单编号列表
    Tids   []int64 `json:"tids,omitempty"`

    // 总数
    Totals   int64 `json:"totals,omitempty"`

}
