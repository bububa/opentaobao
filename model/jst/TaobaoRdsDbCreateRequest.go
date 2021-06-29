package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rds创建数据库 API请求
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

// 初始化TaobaoRdsDbCreateRequest对象
func NewTaobaoRdsDbCreateRequest() *TaobaoRdsDbCreateRequest{
    return &TaobaoRdsDbCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbCreateRequest) GetApiMethodName() string {
    return "taobao.rds.db.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InstanceName Setter
// rds的实例名
func (r *TaobaoRdsDbCreateRequest) SetInstanceName(instanceName string) error {
    r.instanceName = instanceName
    r.Set("instance_name", instanceName)
    return nil
}

// InstanceName Getter
func (r TaobaoRdsDbCreateRequest) GetInstanceName() string {
    return r.instanceName
}
// DbName Setter
// 数据库名
func (r *TaobaoRdsDbCreateRequest) SetDbName(dbName string) error {
    r.dbName = dbName
    r.Set("db_name", dbName)
    return nil
}

// DbName Getter
func (r TaobaoRdsDbCreateRequest) GetDbName() string {
    return r.dbName
}
// AccountName Setter
// 已存在账号名
func (r *TaobaoRdsDbCreateRequest) SetAccountName(accountName string) error {
    r.accountName = accountName
    r.Set("account_name", accountName)
    return nil
}

// AccountName Getter
func (r TaobaoRdsDbCreateRequest) GetAccountName() string {
    return r.accountName
}
