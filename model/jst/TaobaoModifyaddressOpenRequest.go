package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝自助修改地址服务开通 APIRequest
taobao.modifyaddress.open

商家自助修改地址服务开通
*/
type TaobaoModifyaddressOpenRequest struct {
    model.Params

}

func NewTaobaoModifyaddressOpenRequest() *TaobaoModifyaddressOpenRequest{
    return &TaobaoModifyaddressOpenRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoModifyaddressOpenRequest) GetApiMethodName() string {
    return "taobao.modifyaddress.open"
}

func (r TaobaoModifyaddressOpenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


