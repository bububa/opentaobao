package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordsdbgetdbAPIRequest rds获取RDS的DB API请求
// taobao.rds.db.getdb
//
// rds获取RDS的DB
type TaobaordsdbgetdbAPIRequest struct {
	model.Params
	// 账户名
	_accountName string
	// 实例名
	_instanceName string
}

// NewTaobaordsdbgetdbRequest 初始化TaobaordsdbgetdbAPIRequest对象
func NewTaobaordsdbgetdbRequest() *TaobaordsdbgetdbAPIRequest {
	return &TaobaordsdbgetdbAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordsdbgetdbAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.getdb"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordsdbgetdbAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordsdbgetdbAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountName is AccountName Setter
// 账户名
func (r *TaobaordsdbgetdbAPIRequest) SetAccountName(_accountName string) error {
	r._accountName = _accountName
	r.Set("account_name", _accountName)
	return nil
}

// GetAccountName AccountName Getter
func (r TaobaordsdbgetdbAPIRequest) GetAccountName() string {
	return r._accountName
}

// SetInstanceName is InstanceName Setter
// 实例名
func (r *TaobaordsdbgetdbAPIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("instance_name", _instanceName)
	return nil
}

// GetInstanceName InstanceName Getter
func (r TaobaordsdbgetdbAPIRequest) GetInstanceName() string {
	return r._instanceName
}
