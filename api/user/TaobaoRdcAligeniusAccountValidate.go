package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
AG商家账号校验 
taobao.rdc.aligenius.account.validate

提供应对接AG的erp系统查询其旗下的商家是否为AG商家
*/
func TaobaoRdcAligeniusAccountValidate(clt *core.SDKClient, req *user.TaobaoRdcAligeniusAccountValidateRequest, session string) (*user.TaobaoRdcAligeniusAccountValidateAPIResponse, error) {
    var resp user.TaobaoRdcAligeniusAccountValidateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
