package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单确认接口 API请求
taobao.qimen.deliveryorder.confirm

taobao.qimen.deliveryorder.confirm
*/
type TaobaoQimenDeliveryorderConfirmRequest struct {
    model.Params
    // 
    request   *DeliveryOrderConfirmRequest
}

// 初始化TaobaoQimenDeliveryorderConfirmRequest对象
func NewTaobaoQimenDeliveryorderConfirmRequest() *TaobaoQimenDeliveryorderConfirmRequest{
    return &TaobaoQimenDeliveryorderConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderConfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenDeliveryorderConfirmRequest) SetRequest(request *DeliveryOrderConfirmRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenDeliveryorderConfirmRequest) GetRequest() *DeliveryOrderConfirmRequest {
    return r.request
}
