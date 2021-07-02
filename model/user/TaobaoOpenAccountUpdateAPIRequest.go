package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountUpdateAPIRequest Open Account数据更新 API请求
// taobao.open.account.update
//
// Open Account数据更新
type TaobaoOpenAccountUpdateAPIRequest struct {
	model.Params
	// Open Account
	_paramList []OpenAccount
}

// NewTaobaoOpenAccountUpdateRequest 初始化TaobaoOpenAccountUpdateAPIRequest对象
func NewTaobaoOpenAccountUpdateRequest() *TaobaoOpenAccountUpdateAPIRequest {
	return &TaobaoOpenAccountUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamList is ParamList Setter
// Open Account
func (r *TaobaoOpenAccountUpdateAPIRequest) SetParamList(_paramList []OpenAccount) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TaobaoOpenAccountUpdateAPIRequest) GetParamList() []OpenAccount {
	return r._paramList
}
