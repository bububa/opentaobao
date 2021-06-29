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
type TaobaoQimenDeliveryorderCreateRequest struct {
    model.Params
    // 
    request   *DeliveryOrderCreateRequest
}

// 初始化TaobaoQimenDeliveryorderCreateRequest对象
func NewTaobaoQimenDeliveryorderCreateRequest() *TaobaoQimenDeliveryorderCreateRequest{
    return &TaobaoQimenDeliveryorderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenDeliveryorderCreateRequest) SetRequest(request *DeliveryOrderCreateRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenDeliveryorderCreateRequest) GetRequest() *DeliveryOrderCreateRequest {
    return r.request
}
