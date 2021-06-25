package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据收件人信息查询交易单号接口 APIRequest
taobao.qimen.order.query

WMS 调用该接口，根据收件人信息查询平台交易订单号。
*/
type TaobaoQimenOrderQueryRequest struct {
    model.Params

    // request
    request   *Request 

}

func NewTaobaoQimenOrderQueryRequest() *TaobaoQimenOrderQueryRequest{
    return &TaobaoQimenOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenOrderQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.order.query"
}

func (r TaobaoQimenOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenOrderQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenOrderQueryRequest) GetRequest() *Request {
    return r.request
}

