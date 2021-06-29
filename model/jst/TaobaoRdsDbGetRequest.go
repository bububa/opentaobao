package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询rds下的数据库 API请求
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

// 初始化TaobaoRdsDbGetRequest对象
func NewTaobaoRdsDbGetRequest() *TaobaoRdsDbGetRequest{
    return &TaobaoRdsDbGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbGetRequest) GetApiMethodName() string {
    return "taobao.rds.db.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InstanceName Setter
// rds的实例名
func (r *TaobaoRdsDbGetRequest) SetInstanceName(instanceName string) error {
    r.instanceName = instanceName
    r.Set("instance_name", instanceName)
    return nil
}

// InstanceName Getter
func (r TaobaoRdsDbGetRequest) GetInstanceName() string {
    return r.instanceName
}
// DbStatus Setter
// 数据库状态，默认值1
func (r *TaobaoRdsDbGetRequest) SetDbStatus(dbStatus int64) error {
    r.dbStatus = dbStatus
    r.Set("db_status", dbStatus)
    return nil
}

// DbStatus Getter
func (r TaobaoRdsDbGetRequest) GetDbStatus() int64 {
    return r.dbStatus
}
