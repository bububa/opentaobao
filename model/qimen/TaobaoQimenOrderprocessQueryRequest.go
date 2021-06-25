package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单流水查询接口 APIRequest
taobao.qimen.orderprocess.query

ERP调用订单流水查询接口
*/
type TaobaoQimenOrderprocessQueryRequest struct {
    model.Params

    // 
    request   *OrderProcessQueryRequest 

}

func NewTaobaoQimenOrderprocessQueryRequest() *TaobaoQimenOrderprocessQueryRequest{
    return &TaobaoQimenOrderprocessQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenOrderprocessQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.orderprocess.query"
}

func (r TaobaoQimenOrderprocessQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenOrderprocessQueryRequest) SetRequest(request *OrderProcessQueryRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenOrderprocessQueryRequest) GetRequest() *OrderProcessQueryRequest {
    return r.request
}

