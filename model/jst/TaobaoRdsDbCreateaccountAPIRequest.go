package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdsDbCreateaccountAPIRequest
rds创建数据库账户 API请求
taobao.rds.db.createaccount

rds创建数据库账户 */
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
func (r TaobaoRdsDbCreateaccountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 入参对象
func (r *TaobaoRdsDbCreateaccountAPIRequest) SetParam0(_param0 *RequestDbAccountModel) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoRdsDbCreateaccountAPIRequest) GetParam0() *RequestDbAccountModel {
	return r._param0
}
