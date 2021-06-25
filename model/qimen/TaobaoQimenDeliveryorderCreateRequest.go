package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建接口 APIRequest
taobao.qimen.deliveryorder.create

taobao.qimen.deliveryorder.create
*/
type TaobaoQimenDeliveryorderCreateRequest struct {
    model.Params

    // 
    request   *DeliveryOrderCreateRequest 

}

func NewTaobaoQimenDeliveryorderCreateRequest() *TaobaoQimenDeliveryorderCreateRequest{
    return &TaobaoQimenDeliveryorderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenDeliveryorderCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.create"
}

func (r TaobaoQimenDeliveryorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenDeliveryorderCreateRequest) SetRequest(request *DeliveryOrderCreateRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenDeliveryorderCreateRequest) GetRequest() *DeliveryOrderCreateRequest {
    return r.request
}

