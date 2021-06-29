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
type TaobaoQimenOrderprocessQueryRequest struct {
    model.Params
    // 
    request   *OrderProcessQueryRequest
}

// 初始化TaobaoQimenOrderprocessQueryRequest对象
func NewTaobaoQimenOrderprocessQueryRequest() *TaobaoQimenOrderprocessQueryRequest{
    return &TaobaoQimenOrderprocessQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderprocessQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.orderprocess.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderprocessQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenOrderprocessQueryRequest) SetRequest(request *OrderProcessQueryRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenOrderprocessQueryRequest) GetRequest() *OrderProcessQueryRequest {
    return r.request
}
