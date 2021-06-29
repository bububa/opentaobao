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
type TaobaoRdsDbGetdbRequest struct {
    model.Params
    // 账户名
    _accountName   string
    // 实例名
    _instanceName   string
}

// 初始化TaobaoRdsDbGetdbRequest对象
func NewTaobaoRdsDbGetdbRequest() *TaobaoRdsDbGetdbRequest{
    return &TaobaoRdsDbGetdbRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbGetdbRequest) GetApiMethodName() string {
    return "taobao.rds.db.getdb"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbGetdbRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountName Setter
// 账户名
func (r *TaobaoRdsDbGetdbRequest) SetAccountName(_accountName string) error {
    r._accountName = _accountName
    r.Set("account_name", _accountName)
    return nil
}

// AccountName Getter
func (r TaobaoRdsDbGetdbRequest) GetAccountName() string {
    return r._accountName
}
// InstanceName Setter
// 实例名
func (r *TaobaoRdsDbGetdbRequest) SetInstanceName(_instanceName string) error {
    r._instanceName = _instanceName
    r.Set("instance_name", _instanceName)
    return nil
}

// InstanceName Getter
func (r TaobaoRdsDbGetdbRequest) GetInstanceName() string {
    return r._instanceName
}
