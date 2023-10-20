package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenaccountupdateAPIRequest Open Account数据更新 API请求
// taobao.open.account.update
//
// Open Account数据更新
type TaobaoopenaccountupdateAPIRequest struct {
	model.Params
	// Open Account
	_paramList []OpenAccount
}

// NewTaobaoopenaccountupdateRequest 初始化TaobaoopenaccountupdateAPIRequest对象
func NewTaobaoopenaccountupdateRequest() *TaobaoopenaccountupdateAPIRequest {
	return &TaobaoopenaccountupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenaccountupdateAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenaccountupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenaccountupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// Open Account
func (r *TaobaoopenaccountupdateAPIRequest) SetParamList(_paramList []OpenAccount) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TaobaoopenaccountupdateAPIRequest) GetParamList() []OpenAccount {
	return r._paramList
}
