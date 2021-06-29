package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自送根据核销码查订单 API请求
taobao.omniorder.dtd.query

门店自送根据核销码码查询订单信息
*/
type TaobaoOmniorderDtdQueryRequest struct {
    model.Params
    // 核销码
    _code   string
}

// 初始化TaobaoOmniorderDtdQueryRequest对象
func NewTaobaoOmniorderDtdQueryRequest() *TaobaoOmniorderDtdQueryRequest{
    return &TaobaoOmniorderDtdQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderDtdQueryRequest) GetApiMethodName() string {
    return "taobao.omniorder.dtd.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderDtdQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 核销码
func (r *TaobaoOmniorderDtdQueryRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoOmniorderDtdQueryRequest) GetCode() string {
    return r._code
}
