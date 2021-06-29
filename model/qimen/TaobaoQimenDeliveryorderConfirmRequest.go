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
    _request   *DeliveryOrderConfirmRequest
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
func (r *TaobaoQimenDeliveryorderConfirmRequest) SetRequest(_request *DeliveryOrderConfirmRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenDeliveryorderConfirmRequest) GetRequest() *DeliveryOrderConfirmRequest {
    return r._request
}
