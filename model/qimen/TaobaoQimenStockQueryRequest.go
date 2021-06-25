package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存查询接口（多条件） APIRequest
taobao.qimen.stock.query

ERP调用奇门的接口,查询商品的库存量
*/
type TaobaoQimenStockQueryRequest struct {
    model.Params

    // 
    request   *StockQueryRequest 

}

func NewTaobaoQimenStockQueryRequest() *TaobaoQimenStockQueryRequest{
    return &TaobaoQimenStockQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStockQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.stock.query"
}

func (r TaobaoQimenStockQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStockQueryRequest) SetRequest(request *StockQueryRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenStockQueryRequest) GetRequest() *StockQueryRequest {
    return r.request
}

