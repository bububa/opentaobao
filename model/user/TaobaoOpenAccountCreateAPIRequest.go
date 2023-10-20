package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenaccountcreateAPIRequest Open Account导入数据 API请求
// taobao.open.account.create
//
// Open Account导入数据
type TaobaoopenaccountcreateAPIRequest struct {
	model.Params
	// Open Account的列表
	_paramList []OpenAccount
}

// NewTaobaoopenaccountcreateRequest 初始化TaobaoopenaccountcreateAPIRequest对象
func NewTaobaoopenaccountcreateRequest() *TaobaoopenaccountcreateAPIRequest {
	return &TaobaoopenaccountcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenaccountcreateAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenaccountcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenaccountcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// Open Account的列表
func (r *TaobaoopenaccountcreateAPIRequest) SetParamList(_paramList []OpenAccount) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TaobaoopenaccountcreateAPIRequest) GetParamList() []OpenAccount {
	return r._paramList
}
