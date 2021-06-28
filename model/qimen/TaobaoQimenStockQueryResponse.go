package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
库存查询接口（多条件） APIResponse
taobao.qimen.stock.query

ERP调用奇门的接口,查询商品的库存量
*/
type TaobaoQimenStockQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_stock_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *StockQueryResponse `json:"response,omitempty" xml:"