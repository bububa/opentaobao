package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdsDbGetAPIRequest 查询rds下的数据库 API请求
// taobao.rds.db.get
//
// 查询rds实例下的数据库
type TaobaoRdsDbGetAPIRequest struct {
	model.Params
	// rds的实例名
	_instanceName string
	// 数据库状态，默认值1
	_dbStatus int64
}

// NewTaobaoRdsDbGetRequest 初始化TaobaoRdsDbGetAPIRequest对象
func NewTaobaoRdsDbGetRequest() *TaobaoRdsDbGetAPIRequest {
	return &TaobaoRdsDbGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdsDbGetAPIRequest) Reset() {
	r._instanceName = ""
	r._dbStatus = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbGetAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdsDbGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstanceName is InstanceName Setter
// rds的实例名
func (r *TaobaoRdsDbGetAPIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("instance_name", _instanceName)
	return nil
}

// GetInstanceName InstanceName Getter
func (r TaobaoRdsDbGetAPIRequest) GetInstanceName() string {
	return r._instanceName
}

// SetDbStatus is DbStatus Setter
// 数据库状态，默认值1
func (r *TaobaoRdsDbGetAPIRequest) SetDbStatus(_dbStatus int64) error {
	r._dbStatus = _dbStatus
	r.Set("db_status", _dbStatus)
	return nil
}

// GetDbStatus DbStatus Getter
func (r TaobaoRdsDbGetAPIRequest) GetDbStatus() int64 {
	return r._dbStatus
}

var poolTaobaoRdsDbGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdsDbGetRequest()
	},
}

// GetTaobaoRdsDbGetRequest 从 sync.Pool 获取 TaobaoRdsDbGetAPIRequest
func GetTaobaoRdsDbGetAPIRequest() *TaobaoRdsDbGetAPIRequest {
	return poolTaobaoRdsDbGetAPIRequest.Get().(*TaobaoRdsDbGetAPIRequest)
}

// ReleaseTaobaoRdsDbGetAPIRequest 将 TaobaoRdsDbGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdsDbGetAPIRequest(v *TaobaoRdsDbGetAPIRequest) {
	v.Reset()
	poolTaobaoRdsDbGetAPIRequest.Put(v)
}
