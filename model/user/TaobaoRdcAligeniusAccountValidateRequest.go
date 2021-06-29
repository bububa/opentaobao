package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AG商家账号校验 API请求
taobao.rdc.aligenius.account.validate

提供应对接AG的erp系统查询其旗下的商家是否为AG商家
*/
type TaobaoRdcAligeniusAccountValidateRequest struct {
    model.Params
}

// 初始化TaobaoRdcAligeniusAccountValidateRequest对象
func NewTaobaoRdcAligeniusAccountValidateRequest() *TaobaoRdcAligeniusAccountValidateRequest{
    return &TaobaoRdcAligeniusAccountValidateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusAccountValidateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.account.validate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusAccountValidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
