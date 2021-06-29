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
type TaobaoQimenDeliveryorderQueryRequest struct {
    model.Params
    // 
    _request   *DeliveryOrderQueryRequest
}

// 初始化TaobaoQimenDeliveryorderQueryRequest对象
func NewTaobaoQimenDeliveryorderQueryRequest() *TaobaoQimenDeliveryorderQueryRequest{
    return &TaobaoQimenDeliveryorderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenDeliveryorderQueryRequest) SetRequest(_request *DeliveryOrderQueryRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenDeliveryorderQueryRequest) GetRequest() *DeliveryOrderQueryRequest {
    return r._request
}
