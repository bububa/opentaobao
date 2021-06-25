package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rds获取RDS的DB APIRequest
taobao.rds.db.getdb

rds获取RDS的DB
*/
type TaobaoRdsDbGetdbRequest struct {
    model.Params

    // 账户名
    accountName   string 

    // 实例名
    instanceName   string 

}

func NewTaobaoRdsDbGetdbRequest() *TaobaoRdsDbGetdbRequest{
    return &TaobaoRdsDbGetdbRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdsDbGetdbRequest) GetApiMethodName() string {
    return "taobao.rds.db.getdb"
}

func (r TaobaoRdsDbGetdbRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdsDbGetdbRequest) SetAccountName(accountName string) error {
    r.accountName = accountName
    r.Set("account_name", accountName)
    return nil
}

func (r TaobaoRdsDbGetdbRequest) GetAccountName() string {
    return r.accountName
}

func (r *TaobaoRdsDbGetdbRequest) SetInstanceName(instanceName string) error {
    r.instanceName = instanceName
    r.Set("instance_name", instanceName)
    return nil
}

func (r TaobaoRdsDbGetdbRequest) GetInstanceName() string {
    return r.instanceName
}

