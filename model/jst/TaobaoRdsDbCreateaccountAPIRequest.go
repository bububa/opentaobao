package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdsDbCreateaccountAPIRequest rds创建数据库账户 API请求
// taobao.rds.db.createaccount
//
// rds创建数据库账户
type TaobaoRdsDbCreateaccountAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *RequestDbAccountModel
}

// NewTaobaoRdsDbCreateaccountRequest 初始化TaobaoRdsDbCreateaccountAPIRequest对象
func NewTaobaoRdsDbCreateaccountRequest() *TaobaoRdsDbCreateaccountAPIRequest {
	return &TaobaoRdsDbCreateaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbCreateaccountAPIRequest) GetApiMethodName() string {
	return "taobao.rds.db.createaccount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbCreateaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdsDbCreateaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参对象
func (r *TaobaoRdsDbCreateaccountAPIRequest) SetParam0(_param0 *RequestDbAccountModel) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoRdsDbCreateaccountAPIRequest) GetParam0() *RequestDbAccountModel {
	return r._param0
}
