package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建结果通知接口(批量) API请求
taobao.qimen.deliveryorder.batchcreate.answer

WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
*/
type TaobaoQimenDeliveryorderBatchcreateAnswerRequest struct {
    model.Params
    // 
    request   *DeliveryOrderBatchCreateAnswerRequest
}

// 初始化TaobaoQimenDeliveryorderBatchcreateAnswerRequest对象
func NewTaobaoQimenDeliveryorderBatchcreateAnswerRequest() *TaobaoQimenDeliveryorderBatchcreateAnswerRequest{
    return &TaobaoQimenDeliveryorderBatchcreateAnswerRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderBatchcreateAnswerRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.batchcreate.answer"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderBatchcreateAnswerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenDeliveryorderBatchcreateAnswerRequest) SetRequest(request *DeliveryOrderBatchCreateAnswerRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenDeliveryorderBatchcreateAnswerRequest) GetRequest() *DeliveryOrderBatchCreateAnswerRequest {
    return r.request
}
