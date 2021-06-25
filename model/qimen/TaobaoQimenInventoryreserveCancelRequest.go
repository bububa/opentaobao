package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存预占取消接口 APIRequest
taobao.qimen.inventoryreserve.cancel

库存预占取消
*/
type TaobaoQimenInventoryreserveCancelRequest struct {
    model.Params

    // 
    request   *Request 

}

func NewTaobaoQimenInventoryreserveCancelRequest() *TaobaoQimenInventoryreserveCancelRequest{
    return &TaobaoQimenInventoryreserveCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenInventoryreserveCancelRequest) GetApiMethodName() string {
    return "taobao.qimen.inventoryreserve.cancel"
}

func (r TaobaoQimenInventoryreserveCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenInventoryreserveCancelRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenInventoryreserveCancelRequest) GetRequest() *Request {
    return r.request
}

