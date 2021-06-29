package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
库存查询接口（多条件） API返回值 
taobao.qimen.stock.query

ERP调用奇门的接口,查询商品的库存量
*/
type TaobaoQimenStockQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenStockQueryResponse
}

// 库存查询接口（多条件） 成功返回结果
type TaobaoQimenStockQueryResponse struct {
    XMLName xml.Name `xml:"qimen_stock_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *StockQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
