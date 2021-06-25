package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单确认接口 APIRequest
taobao.qimen.deliveryorder.confirm

taobao.qimen.deliveryorder.confirm
*/
type TaobaoQimenDeliveryorderConfirmRequest struct {
    model.Params

    // 
    request   *DeliveryOrderConfirmRequest 

}

func NewTaobaoQimenDeliveryorderConfirmRequest() *TaobaoQimenDeliveryorderConfirmRequest{
    return &TaobaoQimenDeliveryorderConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenDeliveryorderConfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.confirm"
}

func (r TaobaoQimenDeliveryorderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenDeliveryorderConfirmRequest) SetRequest(request *DeliveryOrderConfirmRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenDeliveryorderConfirmRequest) GetRequest() *DeliveryOrderConfirmRequest {
    return r.request
}

