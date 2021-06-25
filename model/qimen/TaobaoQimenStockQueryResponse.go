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
    Response *TaobaoQimenStockQueryResponse `json:"taobao_qimen_stock_query_response,omitempty"`
}

type TaobaoQimenStockQueryResponse struct {

    // 
    Response   *StockQueryResponse `json:"response,omitempty"`

}
