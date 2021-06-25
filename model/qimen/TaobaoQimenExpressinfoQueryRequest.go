package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
配送公司信息查询接口 APIRequest
taobao.qimen.expressinfo.query

配送公司信息查询
*/
type TaobaoQimenExpressinfoQueryRequest struct {
    model.Params

    // 
    request   *Request 

}

func NewTaobaoQimenExpressinfoQueryRequest() *TaobaoQimenExpressinfoQueryRequest{
    return &TaobaoQimenExpressinfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenExpressinfoQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.expressinfo.query"
}

func (r TaobaoQimenExpressinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenExpressinfoQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenExpressinfoQueryRequest) GetRequest() *Request {
    return r.request
}

