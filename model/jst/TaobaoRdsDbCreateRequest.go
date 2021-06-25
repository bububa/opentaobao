package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rds创建数据库 APIRequest
taobao.rds.db.create

在rds实例里创建数据库
*/
type TaobaoRdsDbCreateRequest struct {
    model.Params

    // rds的实例名
    instanceName   string 

    // 数据库名
    dbName   string 

    // 已存在账号名
    accountName   string 

}

func NewTaobaoRdsDbCreateRequest() *TaobaoRdsDbCreateRequest{
    return &TaobaoRdsDbCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdsDbCreateRequest) GetApiMethodName() string {
    return "taobao.rds.db.create"
}

func (r TaobaoRdsDbCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdsDbCreateRequest) SetInstanceName(instanceName string) error {
    r.instanceName = instanceName
    r.Set("instance_name", instanceName)
    return nil
}

func (r TaobaoRdsDbCreateRequest) GetInstanceName() string {
    return r.instanceName
}

func (r *TaobaoRdsDbCreateRequest) SetDbName(dbName string) error {
    r.dbName = dbName
    r.Set("db_name", dbName)
    return nil
}

func (r TaobaoRdsDbCreateRequest) GetDbName() string {
    return r.dbName
}

func (r *TaobaoRdsDbCreateRequest) SetAccountName(accountName string) error {
    r.accountName = accountName
    r.Set("account_name", accountName)
    return nil
}

func (r TaobaoRdsDbCreateRequest) GetAccountName() string {
    return r.accountName
}

