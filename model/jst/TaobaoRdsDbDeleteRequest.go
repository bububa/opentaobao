package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
RDS数据库删除 API请求
taobao.rds.db.delete

通过api删除用户RDS的数据库
*/
type TaobaoRdsDbDeleteRequest struct {
    model.Params
    // rds的实例名
    instanceName   string
    // 数据库的name，可以通过 taobao.rds.db.get 获取
    dbName   string
}

// 初始化TaobaoRdsDbDeleteRequest对象
func NewTaobaoRdsDbDeleteRequest() *TaobaoRdsDbDeleteRequest{
    return &TaobaoRdsDbDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbDeleteRequest) GetApiMethodName() string {
    return "taobao.rds.db.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InstanceName Setter
// rds的实例名
func (r *TaobaoRdsDbDeleteRequest) SetInstanceName(instanceName string) error {
    r.instanceName = instanceName
    r.Set("instance_name", instanceName)
    return nil
}

// InstanceName Getter
func (r TaobaoRdsDbDeleteRequest) GetInstanceName() string {
    return r.instanceName
}
// DbName Setter
// 数据库的name，可以通过 taobao.rds.db.get 获取
func (r *TaobaoRdsDbDeleteRequest) SetDbName(dbName string) error {
    r.dbName = dbName
    r.Set("db_name", dbName)
    return nil
}

// DbName Getter
func (r TaobaoRdsDbDeleteRequest) GetDbName() string {
    return r.dbName
}
