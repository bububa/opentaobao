package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单确认接口 API请求
taobao.qimen.deliveryorder.batchconfirm

taobao.qimen.deliveryorder.batchconfirm
*/
type TaobaoQimenDeliveryorderBatchconfirmRequest struct {
    model.Params
    // 
    request   *DeliveryOrderBatchConfirmRequest
}

// 初始化TaobaoQimenDeliveryorderBatchconfirmRequest对象
func NewTaobaoQimenDeliveryorderBatchconfirmRequest() *TaobaoQimenDeliveryorderBatchconfirmRequest{
    return &TaobaoQimenDeliveryorderBatchconfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderBatchconfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.batchconfirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderBatchconfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenDeliveryorderBatchconfirmRequest) SetRequest(request *DeliveryOrderBatchConfirmRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenDeliveryorderBatchconfirmRequest) GetRequest() *DeliveryOrderBatchConfirmRequest {
    return r.request
}
