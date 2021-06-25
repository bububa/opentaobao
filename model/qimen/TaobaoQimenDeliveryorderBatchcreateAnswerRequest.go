package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建结果通知接口(批量) APIRequest
taobao.qimen.deliveryorder.batchcreate.answer

WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
*/
type TaobaoQimenDeliveryorderBatchcreateAnswerRequest struct {
    model.Params

    // 
    request   *DeliveryOrderBatchCreateAnswerRequest 

}

func NewTaobaoQimenDeliveryorderBatchcreateAnswerRequest() *TaobaoQimenDeliveryorderBatchcreateAnswerRequest{
    return &TaobaoQimenDeliveryorderBatchcreateAnswerRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenDeliveryorderBatchcreateAnswerRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.batchcreate.answer"
}

func (r TaobaoQimenDeliveryorderBatchcreateAnswerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenDeliveryorderBatchcreateAnswerRequest) SetRequest(request *DeliveryOrderBatchCreateAnswerRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenDeliveryorderBatchcreateAnswerRequest) GetRequest() *DeliveryOrderBatchCreateAnswerRequest {
    return r.request
}

