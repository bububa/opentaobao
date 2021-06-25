package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
RDS数据库删除 APIRequest
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

func NewTaobaoRdsDbDeleteRequest() *TaobaoRdsDbDeleteRequest{
    return &TaobaoRdsDbDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdsDbDeleteRequest) GetApiMethodName() string {
    return "taobao.rds.db.delete"
}

func (r TaobaoRdsDbDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdsDbDeleteRequest) SetInstanceName(instanceName string) error {
    r.instanceName = instanceName
    r.Set("instance_name", instanceName)
    return nil
}

func (r TaobaoRdsDbDeleteRequest) GetInstanceName() string {
    return r.instanceName
}

func (r *TaobaoRdsDbDeleteRequest) SetDbName(dbName string) error {
    r.dbName = dbName
    r.Set("db_name", dbName)
    return nil
}

func (r TaobaoRdsDbDeleteRequest) GetDbName() string {
    return r.dbName
}

