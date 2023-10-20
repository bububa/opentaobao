package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordsdbcreateAPIRequest rds创建数据库 API请求
// taobao.rds.db.create
//
// 在rds实例里创建数据库
type TaobaordsdbcreateAPIRequest struct {
	model.Params
	// rds的实例名
	_instanceName string
	// 数据库名
	_dbName string
	// 已存在账号名
	_accountName string
}

// NewTaobaordsdbcreateRequest 初始化TaobaordsdbcreateAPIRequest对象
func NewTaobaordsdbcreateRequest() *TaobaordsdbcreateAPIRequest {
	return &TaobaordsdbcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordsdbcreateAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordsdbcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordsdbcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstanceName is InstanceName Setter
// rds的实例名
func (r *TaobaordsdbcreateAPIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("instance_name", _instanceName)
	return nil
}

// GetInstanceName InstanceName Getter
func (r TaobaordsdbcreateAPIRequest) GetInstanceName() string {
	return r._instanceName
}

// SetDbName is DbName Setter
// 数据库名
func (r *TaobaordsdbcreateAPIRequest) SetDbName(_dbName string) error {
	r._dbName = _dbName
	r.Set("db_name", _dbName)
	return nil
}

// GetDbName DbName Getter
func (r TaobaordsdbcreateAPIRequest) GetDbName() string {
	return r._dbName
}

// SetAccountName is AccountName Setter
// 已存在账号名
func (r *TaobaordsdbcreateAPIRequest) SetAccountName(_accountName string) error {
	r._accountName = _accountName
	r.Set("account_name", _accountName)
	return nil
}

// GetAccountName AccountName Getter
func (r TaobaordsdbcreateAPIRequest) GetAccountName() string {
	return r._accountName
}
