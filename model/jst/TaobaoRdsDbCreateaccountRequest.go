package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rds创建数据库账户 APIRequest
taobao.rds.db.createaccount

rds创建数据库账户
*/
type TaobaoRdsDbCreateaccountRequest struct {
    model.Params

    // 入参对象
    param0   *RequestDbAccountModel 

}

func NewTaobaoRdsDbCreateaccountRequest() *TaobaoRdsDbCreateaccountRequest{
    return &TaobaoRdsDbCreateaccountRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdsDbCreateaccountRequest) GetApiMethodName() string {
    return "taobao.rds.db.createaccount"
}

func (r TaobaoRdsDbCreateaccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdsDbCreateaccountRequest) SetParam0(param0 *RequestDbAccountModel) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoRdsDbCreateaccountRequest) GetParam0() *RequestDbAccountModel {
    return r.param0
}

