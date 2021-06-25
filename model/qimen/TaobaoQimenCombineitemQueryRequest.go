package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合货品关系查询接口 APIRequest
taobao.qimen.combineitem.query

组合货品关系查询
*/
type TaobaoQimenCombineitemQueryRequest struct {
    model.Params

    // 
    request   *Request 

}

func NewTaobaoQimenCombineitemQueryRequest() *TaobaoQimenCombineitemQueryRequest{
    return &TaobaoQimenCombineitemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenCombineitemQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.combineitem.query"
}

func (r TaobaoQimenCombineitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenCombineitemQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenCombineitemQueryRequest) GetRequest() *Request {
    return r.request
}

