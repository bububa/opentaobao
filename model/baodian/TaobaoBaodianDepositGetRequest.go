package baodian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
宝点用户帐户查询（已迁移） API请求
taobao.baodian.deposit.get

查询用户宝点帐户信息及当前宝点价格
*/
type TaobaoBaodianDepositGetRequest struct {
    model.Params
}

// 初始化TaobaoBaodianDepositGetRequest对象
func NewTaobaoBaodianDepositGetRequest() *TaobaoBaodianDepositGetRequest{
    return &TaobaoBaodianDepositGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaodianDepositGetRequest) GetApiMethodName() string {
    return "taobao.baodian.deposit.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaodianDepositGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
