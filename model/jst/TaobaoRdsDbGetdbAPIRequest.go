package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdsDbGetdbAPIRequest rds获取RDS的DB API请求
// taobao.rds.db.getdb
//
// rds获取RDS的DB
type TaobaoRdsDbGetdbAPIRequest struct {
	model.Params
	// 账户名
	_accountName string
	// 实例名
	_instanceName string
}

// NewTaobaoRdsDbGetdbRequest 初始化TaobaoRdsDbGetdbAPIRequest对象
func NewTaobaoRdsDbGetdbRequest() *TaobaoRdsDbGetdbAPIRequest {
	return &TaobaoRdsDbGetdbAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdsDbGetdbAPIRequest) Reset() {
	r._accountName = ""
	r._instanceName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbGetdbAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.getdb"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbGetdbAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdsDbGetdbAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountName is AccountName Setter
// 账户名
func (r *TaobaoRdsDbGetdbAPIRequest) SetAccountName(_accountName string) error {
	r._accountName = _accountName
	r.Set("account_name", _accountName)
	return nil
}

// GetAccountName AccountName Getter
func (r TaobaoRdsDbGetdbAPIRequest) GetAccountName() string {
	return r._accountName
}

// SetInstanceName is InstanceName Setter
// 实例名
func (r *TaobaoRdsDbGetdbAPIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("instance_name", _instanceName)
	return nil
}

// GetInstanceName InstanceName Getter
func (r TaobaoRdsDbGetdbAPIRequest) GetInstanceName() string {
	return r._instanceName
}

var poolTaobaoRdsDbGetdbAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdsDbGetdbRequest()
	},
}

// GetTaobaoRdsDbGetdbRequest 从 sync.Pool 获取 TaobaoRdsDbGetdbAPIRequest
func GetTaobaoRdsDbGetdbAPIRequest() *TaobaoRdsDbGetdbAPIRequest {
	return poolTaobaoRdsDbGetdbAPIRequest.Get().(*TaobaoRdsDbGetdbAPIRequest)
}

// ReleaseTaobaoRdsDbGetdbAPIRequest 将 TaobaoRdsDbGetdbAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdsDbGetdbAPIRequest(v *TaobaoRdsDbGetdbAPIRequest) {
	v.Reset()
	poolTaobaoRdsDbGetdbAPIRequest.Put(v)
}
