package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询rds下的数据库 APIRequest
taobao.rds.db.get

查询rds实例下的数据库
*/
type TaobaoRdsDbGetRequest struct {
    model.Params

    // rds的实例名
    instanceName   string 

    // 数据库状态，默认值1
    dbStatus   int64 

}

func NewTaobaoRdsDbGetRequest() *TaobaoRdsDbGetRequest{
    return &TaobaoRdsDbGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdsDbGetRequest) GetApiMethodName() string {
    return "taobao.rds.db.get"
}

func (r TaobaoRdsDbGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdsDbGetRequest) SetInstanceName(instanceName string) error {
    r.instanceName = instanceName
    r.Set("instance_name", instanceName)
    return nil
}

func (r TaobaoRdsDbGetRequest) GetInstanceName() string {
    return r.instanceName
}

func (r *TaobaoRdsDbGetRequest) SetDbStatus(dbStatus int64) error {
    r.dbStatus = dbStatus
    r.Set("db_status", dbStatus)
    return nil
}

func (r TaobaoRdsDbGetRequest) GetDbStatus() int64 {
    return r.dbStatus
}

