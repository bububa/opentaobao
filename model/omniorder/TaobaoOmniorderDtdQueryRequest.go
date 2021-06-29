package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自送根据核销码查订单 APIRequest
taobao.omniorder.dtd.query

门店自送根据核销码码查询订单信息
*/
type TaobaoOmniorderDtdQueryRequest struct {
    model.Params

    // 核销码
    code   string 

}

func NewTaobaoOmniorderDtdQueryRequest() *TaobaoOmniorderDtdQueryRequest{
    return &TaobaoOmniorderDtdQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderDtdQueryRequest) GetApiMethodName() string {
    return "taobao.omniorder.dtd.query"
}

func (r TaobaoOmniorderDtdQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderDtdQueryRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoOmniorderDtdQueryRequest) GetCode() string {
    return r.code
}

