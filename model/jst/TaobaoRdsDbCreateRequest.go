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
    _instanceName   string
    // 数据库名
    _dbName   string
    // 已存在账号名
    _accountName   string
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
func (r *TaobaoRdsDbCreateRequest) SetInstanceName(_instanceName string) error {
    r._instanceName = _instanceName
    r.Set("instance_name", _instanceName)
    return nil
}

// InstanceName Getter
func (r TaobaoRdsDbCreateRequest) GetInstanceName() string {
    return r._instanceName
}
// DbName Setter
// 数据库名
func (r *TaobaoRdsDbCreateRequest) SetDbName(_dbName string) error {
    r._dbName = _dbName
    r.Set("db_name", _dbName)
    return nil
}

// DbName Getter
func (r TaobaoRdsDbCreateRequest) GetDbName() string {
    return r._dbName
}
// AccountName Setter
// 已存在账号名
func (r *TaobaoRdsDbCreateRequest) SetAccountName(_accountName string) error {
    r._accountName = _accountName
    r.Set("account_name", _accountName)
    return nil
}

// AccountName Getter
func (r TaobaoRdsDbCreateRequest) GetAccountName() string {
    return r._accountName
}
