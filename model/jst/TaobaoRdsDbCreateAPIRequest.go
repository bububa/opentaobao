package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdsDbCreateAPIRequest rds创建数据库 API请求
// taobao.rds.db.create
//
// 在rds实例里创建数据库
type TaobaoRdsDbCreateAPIRequest struct {
	model.Params
	// rds的实例名
	_instanceName string
	// 数据库名
	_dbName string
	// 已存在账号名
	_accountName string
}

// NewTaobaoRdsDbCreateRequest 初始化TaobaoRdsDbCreateAPIRequest对象
func NewTaobaoRdsDbCreateRequest() *TaobaoRdsDbCreateAPIRequest {
	return &TaobaoRdsDbCreateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdsDbCreateAPIRequest) Reset() {
	r._instanceName = ""
	r._dbName = ""
	r._accountName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbCreateAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdsDbCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstanceName is InstanceName Setter
// rds的实例名
func (r *TaobaoRdsDbCreateAPIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("instance_name", _instanceName)
	return nil
}

// GetInstanceName InstanceName Getter
func (r TaobaoRdsDbCreateAPIRequest) GetInstanceName() string {
	return r._instanceName
}

// SetDbName is DbName Setter
// 数据库名
func (r *TaobaoRdsDbCreateAPIRequest) SetDbName(_dbName string) error {
	r._dbName = _dbName
	r.Set("db_name", _dbName)
	return nil
}

// GetDbName DbName Getter
func (r TaobaoRdsDbCreateAPIRequest) GetDbName() string {
	return r._dbName
}

// SetAccountName is AccountName Setter
// 已存在账号名
func (r *TaobaoRdsDbCreateAPIRequest) SetAccountName(_accountName string) error {
	r._accountName = _accountName
	r.Set("account_name", _accountName)
	return nil
}

// GetAccountName AccountName Getter
func (r TaobaoRdsDbCreateAPIRequest) GetAccountName() string {
	return r._accountName
}

var poolTaobaoRdsDbCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdsDbCreateRequest()
	},
}

// GetTaobaoRdsDbCreateRequest 从 sync.Pool 获取 TaobaoRdsDbCreateAPIRequest
func GetTaobaoRdsDbCreateAPIRequest() *TaobaoRdsDbCreateAPIRequest {
	return poolTaobaoRdsDbCreateAPIRequest.Get().(*TaobaoRdsDbCreateAPIRequest)
}

// ReleaseTaobaoRdsDbCreateAPIRequest 将 TaobaoRdsDbCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdsDbCreateAPIRequest(v *TaobaoRdsDbCreateAPIRequest) {
	v.Reset()
	poolTaobaoRdsDbCreateAPIRequest.Put(v)
}
