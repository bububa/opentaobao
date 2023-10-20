package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordsdbgetAPIRequest 查询rds下的数据库 API请求
// taobao.rds.db.get
//
// 查询rds实例下的数据库
type TaobaordsdbgetAPIRequest struct {
	model.Params
	// rds的实例名
	_instanceName string
	// 数据库状态，默认值1
	_dbStatus int64
}

// NewTaobaordsdbgetRequest 初始化TaobaordsdbgetAPIRequest对象
func NewTaobaordsdbgetRequest() *TaobaordsdbgetAPIRequest {
	return &TaobaordsdbgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordsdbgetAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordsdbgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordsdbgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstanceName is InstanceName Setter
// rds的实例名
func (r *TaobaordsdbgetAPIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("instance_name", _instanceName)
	return nil
}

// GetInstanceName InstanceName Getter
func (r TaobaordsdbgetAPIRequest) GetInstanceName() string {
	return r._instanceName
}

// SetDbStatus is DbStatus Setter
// 数据库状态，默认值1
func (r *TaobaordsdbgetAPIRequest) SetDbStatus(_dbStatus int64) error {
	r._dbStatus = _dbStatus
	r.Set("db_status", _dbStatus)
	return nil
}

// GetDbStatus DbStatus Getter
func (r TaobaordsdbgetAPIRequest) GetDbStatus() int64 {
	return r._dbStatus
}
