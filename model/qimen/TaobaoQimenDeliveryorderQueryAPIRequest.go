package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单查询接口 API请求
taobao.qimen.deliveryorder.query

ERP调用奇门的发货单查询接口，查询发货单详情
*/
type TaobaoQimenDeliveryorderQueryAPIRequest struct {
    model.Params
    // 
    _request   *DeliveryOrderQueryRequest
}

// 初始化TaobaoQimenDeliveryorderQueryAPIRequest对象
func NewTaobaoQimenDeliveryorderQueryRequest() *TaobaoQimenDeliveryorderQueryAPIRequest{
    return &TaobaoQimenDeliveryorderQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderQueryAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenDeliveryorderQueryAPIRequest) SetRequest(_request *DeliveryOrderQueryRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenDeliveryorderQueryAPIRequest) GetRequest() *DeliveryOrderQueryRequest {
    return r._request
}
