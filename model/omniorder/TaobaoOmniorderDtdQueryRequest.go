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
type TaobaoOmniorderDtdQueryAPIRequest struct {
    model.Params
    // 核销码
    _code   string
}

// 初始化TaobaoOmniorderDtdQueryAPIRequest对象
func NewTaobaoOmniorderDtdQueryRequest() *TaobaoOmniorderDtdQueryAPIRequest{
    return &TaobaoOmniorderDtdQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderDtdQueryAPIRequest) GetApiMethodName() string {
    return "taobao.omniorder.dtd.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderDtdQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 核销码
func (r *TaobaoOmniorderDtdQueryAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoOmniorderDtdQueryAPIRequest) GetCode() string {
    return r._code
}
