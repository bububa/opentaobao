package baodian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
宝点用户帐户查询（已迁移） APIRequest
taobao.baodian.deposit.get

查询用户宝点帐户信息及当前宝点价格
*/
type TaobaoBaodianDepositGetRequest struct {
    model.Params

}

func NewTaobaoBaodianDepositGetRequest() *TaobaoBaodianDepositGetRequest{
    return &TaobaoBaodianDepositGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaodianDepositGetRequest) GetApiMethodName() string {
    return "taobao.baodian.deposit.get"
}

func (r TaobaoBaodianDepositGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


