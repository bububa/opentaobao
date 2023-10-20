package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordsdbdeleteAPIRequest RDS数据库删除 API请求
// taobao.rds.db.delete
//
// 通过api删除用户RDS的数据库
type TaobaordsdbdeleteAPIRequest struct {
	model.Params
	// rds的实例名
	_instanceName string
	// 数据库的name，可以通过 taobao.rds.db.get 获取
	_dbName string
}

// NewTaobaordsdbdeleteRequest 初始化TaobaordsdbdeleteAPIRequest对象
func NewTaobaordsdbdeleteRequest() *TaobaordsdbdeleteAPIRequest {
	return &TaobaordsdbdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordsdbdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordsdbdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordsdbdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstanceName is InstanceName Setter
// rds的实例名
func (r *TaobaordsdbdeleteAPIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("instance_name", _instanceName)
	return nil
}

// GetInstanceName InstanceName Getter
func (r TaobaordsdbdeleteAPIRequest) GetInstanceName() string {
	return r._instanceName
}

// SetDbName is DbName Setter
// 数据库的name，可以通过 taobao.rds.db.get 获取
func (r *TaobaordsdbdeleteAPIRequest) SetDbName(_dbName string) error {
	r._dbName = _dbName
	r.Set("db_name", _dbName)
	return nil
}

// GetDbName DbName Getter
func (r TaobaordsdbdeleteAPIRequest) GetDbName() string {
	return r._dbName
}
