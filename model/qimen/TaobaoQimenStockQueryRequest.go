package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存查询接口（多条件） API请求
taobao.qimen.stock.query

ERP调用奇门的接口,查询商品的库存量
*/
type TaobaoQimenStockQueryRequest struct {
    model.Params
    // 
    request   *StockQueryRequest
}

// 初始化TaobaoQimenStockQueryRequest对象
func NewTaobaoQimenStockQueryRequest() *TaobaoQimenStockQueryRequest{
    return &TaobaoQimenStockQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStockQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.stock.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStockQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenStockQueryRequest) SetRequest(request *StockQueryRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenStockQueryRequest) GetRequest() *StockQueryRequest {
    return r.request
}
