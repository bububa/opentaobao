package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店自提根据核销码查询订单 APIRequest
taobao.omniorder.storecollect.query

全渠道门店自提根据核销码查询订单
*/
type TaobaoOmniorderStorecollectQueryRequest struct {
    model.Params

    // 核销码
    code   string 

}

func NewTaobaoOmniorderStorecollectQueryRequest() *TaobaoOmniorderStorecollectQueryRequest{
    return &TaobaoOmniorderStorecollectQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStorecollectQueryRequest) GetApiMethodName() string {
    return "taobao.omniorder.storecollect.query"
}

func (r TaobaoOmniorderStorecollectQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStorecollectQueryRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoOmniorderStorecollectQueryRequest) GetCode() string {
    return r.code
}

