package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单确认接口 APIRequest
taobao.qimen.deliveryorder.batchconfirm

taobao.qimen.deliveryorder.batchconfirm
*/
type TaobaoQimenDeliveryorderBatchconfirmRequest struct {
    model.Params

    // 
    request   *DeliveryOrderBatchConfirmRequest 

}

func NewTaobaoQimenDeliveryorderBatchconfirmRequest() *TaobaoQimenDeliveryorderBatchconfirmRequest{
    return &TaobaoQimenDeliveryorderBatchconfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenDeliveryorderBatchconfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.batchconfirm"
}

func (r TaobaoQimenDeliveryorderBatchconfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenDeliveryorderBatchconfirmRequest) SetRequest(request *DeliveryOrderBatchConfirmRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenDeliveryorderBatchconfirmRequest) GetRequest() *DeliveryOrderBatchConfirmRequest {
    return r.request
}

