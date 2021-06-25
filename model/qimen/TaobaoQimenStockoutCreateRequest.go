package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单创建接口 APIRequest
taobao.qimen.stockout.create

ERP调用奇门接口，创建出库单信息
*/
type TaobaoQimenStockoutCreateRequest struct {
    model.Params

    // 
    request   *StockOutCreateRequest 

}

func NewTaobaoQimenStockoutCreateRequest() *TaobaoQimenStockoutCreateRequest{
    return &TaobaoQimenStockoutCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStockoutCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.stockout.create"
}

func (r TaobaoQimenStockoutCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStockoutCreateRequest) SetRequest(request *StockOutCreateRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenStockoutCreateRequest) GetRequest() *StockOutCreateRequest {
    return r.request
}

