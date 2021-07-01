package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单流水查询接口 API请求
taobao.qimen.orderprocess.query

ERP调用订单流水查询接口
*/
type TaobaoQimenOrderprocessQueryAPIRequest struct {
    model.Params
    // 
    _request   *OrderProcessQueryRequest
}

// 初始化TaobaoQimenOrderprocessQueryAPIRequest对象
func NewTaobaoQimenOrderprocessQueryRequest() *TaobaoQimenOrderprocessQueryAPIRequest{
    return &TaobaoQimenOrderprocessQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.orderprocess.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenOrderprocessQueryAPIRequest) SetRequest(_request *OrderProcessQueryRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetRequest() *OrderProcessQueryRequest {
    return r._request
}
