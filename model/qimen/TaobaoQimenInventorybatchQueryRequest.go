package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品单仓批次库存查询接口 APIRequest
taobao.qimen.inventorybatch.query

ERP 通过该接口查询指定商品的单仓批次库存
*/
type TaobaoQimenInventorybatchQueryRequest struct {
    model.Params

    // request
    request   *Request 

}

func NewTaobaoQimenInventorybatchQueryRequest() *TaobaoQimenInventorybatchQueryRequest{
    return &TaobaoQimenInventorybatchQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenInventorybatchQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.inventorybatch.query"
}

func (r TaobaoQimenInventorybatchQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenInventorybatchQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenInventorybatchQueryRequest) GetRequest() *Request {
    return r.request
}

