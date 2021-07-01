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
type TaobaoQimenDeliveryorderConfirmAPIRequest struct {
    model.Params
    // 
    _request   *DeliveryOrderConfirmRequest
}

// 初始化TaobaoQimenDeliveryorderConfirmAPIRequest对象
func NewTaobaoQimenDeliveryorderConfirmRequest() *TaobaoQimenDeliveryorderConfirmAPIRequest{
    return &TaobaoQimenDeliveryorderConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderConfirmAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenDeliveryorderConfirmAPIRequest) SetRequest(_request *DeliveryOrderConfirmRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenDeliveryorderConfirmAPIRequest) GetRequest() *DeliveryOrderConfirmRequest {
    return r._request
}
