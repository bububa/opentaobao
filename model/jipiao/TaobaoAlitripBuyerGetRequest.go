package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
敏感信息查询 API请求
taobao.alitrip.buyer.get

针对商家提供统一的TOP接口，可以根据订单获取订单对应买家联系电话（阿里小号）。
*/
type TaobaoAlitripBuyerGetRequest struct {
    model.Params
    // 敏感信息查询请求参数
    _requestAxb   *RequestAxbDO
}

// 初始化TaobaoAlitripBuyerGetRequest对象
func NewTaobaoAlitripBuyerGetRequest() *TaobaoAlitripBuyerGetRequest{
    return &TaobaoAlitripBuyerGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripBuyerGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.buyer.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripBuyerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestAxb Setter
// 敏感信息查询请求参数
func (r *TaobaoAlitripBuyerGetRequest) SetRequestAxb(_requestAxb *RequestAxbDO) error {
    r._requestAxb = _requestAxb
    r.Set("request_axb", _requestAxb)
    return nil
}

// RequestAxb Getter
func (r TaobaoAlitripBuyerGetRequest) GetRequestAxb() *RequestAxbDO {
    return r._requestAxb
}
