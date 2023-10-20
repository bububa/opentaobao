package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordsdbcreateaccountAPIRequest rds创建数据库账户 API请求
// taobao.rds.db.createaccount
//
// rds创建数据库账户
type TaobaordsdbcreateaccountAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *RequestDbAccountModel
}

// NewTaobaordsdbcreateaccountRequest 初始化TaobaordsdbcreateaccountAPIRequest对象
func NewTaobaordsdbcreateaccountRequest() *TaobaordsdbcreateaccountAPIRequest {
	return &TaobaordsdbcreateaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordsdbcreateaccountAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.createaccount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordsdbcreateaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordsdbcreateaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参对象
func (r *TaobaordsdbcreateaccountAPIRequest) SetParam0(_param0 *RequestDbAccountModel) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaordsdbcreateaccountAPIRequest) GetParam0() *RequestDbAccountModel {
	return r._param0
}
