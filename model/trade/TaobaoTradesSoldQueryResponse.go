package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据收件人信息查询交易单号 APIResponse
taobao.trades.sold.query

根据收件人信息查询交易单号。
*/
type TaobaoTradesSoldQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTradesSoldQueryResponse `json:"trades_sold_query_response,omitempty"` 
    TaobaoTradesSoldQueryResponse
}

/* model for simplify = false
type TaobaoTradesSoldQueryResponse struct {

    // 订单ID列表。按照订单创建时间倒序，最多返回最近的100笔订单。
    
    TidList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"tid_list,omitempty"`
    

}
*/

type TaobaoTradesSoldQueryResponse struct {

    // 订单ID列表。按照订单创建时间倒序，最多返回最近的100笔订单。
    TidList   []string `json:"tid_list,omitempty"`

}
