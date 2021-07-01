package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rds获取RDS的DB API请求
taobao.rds.db.getdb

rds获取RDS的DB
*/
type TaobaoRdsDbGetdbAPIRequest struct {
    model.Params
    // 账户名
    _accountName   string
    // 实例名
    _instanceName   string
}

// 初始化TaobaoRdsDbGetdbAPIRequest对象
func NewTaobaoRdsDbGetdbRequest() *TaobaoRdsDbGetdbAPIRequest{
    return &TaobaoRdsDbGetdbAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbGetdbAPIRequest) GetApiMethodName() string {
    return "taobao.rds.db.getdb"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbGetdbAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountName Setter
// 账户名
func (r *TaobaoRdsDbGetdbAPIRequest) SetAccountName(_accountName string) error {
    r._accountName = _accountName
    r.Set("account_name", _accountName)
    return nil
}

// AccountName Getter
func (r TaobaoRdsDbGetdbAPIRequest) GetAccountName() string {
    return r._accountName
}
// InstanceName Setter
// 实例名
func (r *TaobaoRdsDbGetdbAPIRequest) SetInstanceName(_instanceName string) error {
    r._instanceName = _instanceName
    r.Set("instance_name", _instanceName)
    return nil
}

// InstanceName Getter
func (r TaobaoRdsDbGetdbAPIRequest) GetInstanceName() string {
    return r._instanceName
}
