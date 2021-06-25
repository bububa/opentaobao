package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AG商家账号校验 APIRequest
taobao.rdc.aligenius.account.validate

提供应对接AG的erp系统查询其旗下的商家是否为AG商家
*/
type TaobaoRdcAligeniusAccountValidateRequest struct {
    model.Params

}

func NewTaobaoRdcAligeniusAccountValidateRequest() *TaobaoRdcAligeniusAccountValidateRequest{
    return &TaobaoRdcAligeniusAccountValidateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdcAligeniusAccountValidateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.account.validate"
}

func (r TaobaoRdcAligeniusAccountValidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


