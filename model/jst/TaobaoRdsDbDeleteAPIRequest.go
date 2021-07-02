package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdsDbDeleteAPIRequest RDS数据库删除 API请求
// taobao.rds.db.delete
//
// 通过api删除用户RDS的数据库
type TaobaoRdsDbDeleteAPIRequest struct {
	model.Params
	// rds的实例名
	_instanceName string
	// 数据库的name，可以通过 taobao.rds.db.get 获取
	_dbName string
}

// NewTaobaoRdsDbDeleteRequest 初始化TaobaoRdsDbDeleteAPIRequest对象
func NewTaobaoRdsDbDeleteRequest() *TaobaoRdsDbDeleteAPIRequest {
	return &TaobaoRdsDbDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is InstanceName Setter
// rds的实例名
func (r *TaobaoRdsDbDeleteAPIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("instance_name", _instanceName)
	return nil
}

// Get InstanceName Getter
func (r TaobaoRdsDbDeleteAPIRequest) GetInstanceName() string {
	return r._instanceName
}

// Set is DbName Setter
// 数据库的name，可以通过 taobao.rds.db.get 获取
func (r *TaobaoRdsDbDeleteAPIRequest) SetDbName(_dbName string) error {
	r._dbName = _dbName
	r.Set("db_name", _dbName)
	return nil
}

// Get DbName Getter
func (r TaobaoRdsDbDeleteAPIRequest) GetDbName() string {
	return r._dbName
}
