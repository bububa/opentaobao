package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店自提根据核销码查询订单 API请求
taobao.omniorder.storecollect.query

全渠道门店自提根据核销码查询订单
*/
type TaobaoOmniorderStorecollectQueryRequest struct {
    model.Params
    // 核销码
    _code   string
}

// 初始化TaobaoOmniorderStorecollectQueryRequest对象
func NewTaobaoOmniorderStorecollectQueryRequest() *TaobaoOmniorderStorecollectQueryRequest{
    return &TaobaoOmniorderStorecollectQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStorecollectQueryRequest) GetApiMethodName() string {
    return "taobao.omniorder.storecollect.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStorecollectQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 核销码
func (r *TaobaoOmniorderStorecollectQueryRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoOmniorderStorecollectQueryRequest) GetCode() string {
    return r._code
}
