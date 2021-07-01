package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据收件人信息查询交易单号接口 API请求
taobao.qimen.order.query

WMS 调用该接口，根据收件人信息查询平台交易订单号。
*/
type TaobaoQimenOrderQueryAPIRequest struct {
    model.Params
    // request
    _request   *TaobaoQimenOrderQueryRequest
}

// 初始化TaobaoQimenOrderQueryAPIRequest对象
func NewTaobaoQimenOrderQueryRequest() *TaobaoQimenOrderQueryAPIRequest{
    return &TaobaoQimenOrderQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderQueryAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.order.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// request
func (r *TaobaoQimenOrderQueryAPIRequest) SetRequest(_request *TaobaoQimenOrderQueryRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenOrderQueryAPIRequest) GetRequest() *TaobaoQimenOrderQueryRequest {
    return r._request
}
