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
    _instanceName   string
    // 数据库状态，默认值1
    _dbStatus   int64
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
func (r *TaobaoRdsDbGetRequest) SetInstanceName(_instanceName string) error {
    r._instanceName = _instanceName
    r.Set("instance_name", _instanceName)
    return nil
}

// InstanceName Getter
func (r TaobaoRdsDbGetRequest) GetInstanceName() string {
    return r._instanceName
}
// DbStatus Setter
// 数据库状态，默认值1
func (r *TaobaoRdsDbGetRequest) SetDbStatus(_dbStatus int64) error {
    r._dbStatus = _dbStatus
    r.Set("db_status", _dbStatus)
    return nil
}

// DbStatus Getter
func (r TaobaoRdsDbGetRequest) GetDbStatus() int64 {
    return r._dbStatus
}
