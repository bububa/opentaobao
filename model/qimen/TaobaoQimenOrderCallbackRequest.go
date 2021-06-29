package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
配送拦截接口 API请求
taobao.qimen.order.callback

配送拦截
*/
type TaobaoQimenOrderCallbackRequest struct {
    model.Params
    // 
    _request   *OrderCallbackRequestDO
}

// 初始化TaobaoQimenOrderCallbackRequest对象
func NewTaobaoQimenOrderCallbackRequest() *TaobaoQimenOrderCallbackRequest{
    return &TaobaoQimenOrderCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderCallbackRequest) GetApiMethodName() string {
    return "taobao.qimen.order.callback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenOrderCallbackRequest) SetRequest(_request *OrderCallbackRequestDO) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenOrderCallbackRequest) GetRequest() *OrderCallbackRequestDO {
    return r._request
}
