package jst

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdsDbDeleteAPIRequest) Reset() {
	r._instanceName = ""
	r._dbName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdsDbDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstanceName is InstanceName Setter
// rds的实例名
func (r *TaobaoRdsDbDeleteAPIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("instance_name", _instanceName)
	return nil
}

// GetInstanceName InstanceName Getter
func (r TaobaoRdsDbDeleteAPIRequest) GetInstanceName() string {
	return r._instanceName
}

// SetDbName is DbName Setter
// 数据库的name，可以通过 taobao.rds.db.get 获取
func (r *TaobaoRdsDbDeleteAPIRequest) SetDbName(_dbName string) error {
	r._dbName = _dbName
	r.Set("db_name", _dbName)
	return nil
}

// GetDbName DbName Getter
func (r TaobaoRdsDbDeleteAPIRequest) GetDbName() string {
	return r._dbName
}

var poolTaobaoRdsDbDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdsDbDeleteRequest()
	},
}

// GetTaobaoRdsDbDeleteRequest 从 sync.Pool 获取 TaobaoRdsDbDeleteAPIRequest
func GetTaobaoRdsDbDeleteAPIRequest() *TaobaoRdsDbDeleteAPIRequest {
	return poolTaobaoRdsDbDeleteAPIRequest.Get().(*TaobaoRdsDbDeleteAPIRequest)
}

// ReleaseTaobaoRdsDbDeleteAPIRequest 将 TaobaoRdsDbDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdsDbDeleteAPIRequest(v *TaobaoRdsDbDeleteAPIRequest) {
	v.Reset()
	poolTaobaoRdsDbDeleteAPIRequest.Put(v)
}
