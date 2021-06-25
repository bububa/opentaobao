package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建批量接口 APIRequest
taobao.qimen.deliveryorder.batchcreate

ERP调用接口，将发货信息批量推送给WMS
*/
type TaobaoQimenDeliveryorderBatchcreateRequest struct {
    model.Params

    // 
    request   *DeliveryOrderBatchCreateRequest 

}

func NewTaobaoQimenDeliveryorderBatchcreateRequest() *TaobaoQimenDeliveryorderBatchcreateRequest{
    return &TaobaoQimenDeliveryorderBatchcreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenDeliveryorderBatchcreateRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.batchcreate"
}

func (r TaobaoQimenDeliveryorderBatchcreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenDeliveryorderBatchcreateRequest) SetRequest(request *DeliveryOrderBatchCreateRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenDeliveryorderBatchcreateRequest) GetRequest() *DeliveryOrderBatchCreateRequest {
    return r.request
}

