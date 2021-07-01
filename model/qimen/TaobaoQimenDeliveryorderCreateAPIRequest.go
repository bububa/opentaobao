package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建接口 API请求
taobao.qimen.deliveryorder.create

taobao.qimen.deliveryorder.create
*/
type TaobaoQimenDeliveryorderCreateAPIRequest struct {
    model.Params
    // 
    _request   *DeliveryOrderCreateRequest
}

// 初始化TaobaoQimenDeliveryorderCreateAPIRequest对象
func NewTaobaoQimenDeliveryorderCreateRequest() *TaobaoQimenDeliveryorderCreateAPIRequest{
    return &TaobaoQimenDeliveryorderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderCreateAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenDeliveryorderCreateAPIRequest) SetRequest(_request *DeliveryOrderCreateRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenDeliveryorderCreateAPIRequest) GetRequest() *DeliveryOrderCreateRequest {
    return r._request
}
