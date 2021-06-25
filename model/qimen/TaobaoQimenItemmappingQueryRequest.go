package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
前后端商品映射查询接口 APIRequest
taobao.qimen.itemmapping.query

前后端商品映射查询接口
*/
type TaobaoQimenItemmappingQueryRequest struct {
    model.Params

    // 
    request   *Request 

}

func NewTaobaoQimenItemmappingQueryRequest() *TaobaoQimenItemmappingQueryRequest{
    return &TaobaoQimenItemmappingQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenItemmappingQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.itemmapping.query"
}

func (r TaobaoQimenItemmappingQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenItemmappingQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenItemmappingQueryRequest) GetRequest() *Request {
    return r.request
}

