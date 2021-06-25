package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存查询接口（多商品） APIRequest
taobao.qimen.inventory.query

ERP调用奇门的接口,查询商品的库存量
*/
type TaobaoQimenInventoryQueryRequest struct {
    model.Params

    // 
    request   *InventoryQueryRequest 

}

func NewTaobaoQimenInventoryQueryRequest() *TaobaoQimenInventoryQueryRequest{
    return &TaobaoQimenInventoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenInventoryQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.inventory.query"
}

func (r TaobaoQimenInventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenInventoryQueryRequest) SetRequest(request *InventoryQueryRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenInventoryQueryRequest) GetRequest() *InventoryQueryRequest {
    return r.request
}

