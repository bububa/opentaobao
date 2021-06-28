package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
库存查询接口（多条件） APIResponse
taobao.qimen.stock.query

ERP调用奇门的接口,查询商品的库存量
*/
type TaobaoQimenStockQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenStockQueryResponse `json:"qimen_stock_query_response,omitempty"` 
    TaobaoQimenStockQueryResponse
}

/* model for simplify = false
type TaobaoQimenStockQueryResponse struct {

    // 
    
    Response  *struct {
        StockQueryResponse  *StockQueryResponse `json:"stock_query_response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenStockQueryResponse struct {

    // 
    Response   *StockQueryResponse `json:"response,omitempty"`

}
