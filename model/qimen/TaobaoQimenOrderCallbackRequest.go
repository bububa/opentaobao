package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
配送拦截接口 APIRequest
taobao.qimen.order.callback

配送拦截
*/
type TaobaoQimenOrderCallbackRequest struct {
    model.Params

    // 
    request   *OrderCallbackRequestDO 

}

func NewTaobaoQimenOrderCallbackRequest() *TaobaoQimenOrderCallbackRequest{
    return &TaobaoQimenOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenOrderCallbackRequest) GetApiMethodName() string {
    return "taobao.qimen.order.callback"
}

func (r TaobaoQimenOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenOrderCallbackRequest) SetRequest(request *OrderCallbackRequestDO) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenOrderCallbackRequest) GetRequest() *OrderCallbackRequestDO {
    return r.request
}

