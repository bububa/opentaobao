package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountCreateAPIRequest Open Account导入数据 API请求
// taobao.open.account.create
//
// Open Account导入数据
type TaobaoOpenAccountCreateAPIRequest struct {
	model.Params
	// Open Account的列表
	_paramList []OpenAccount
}

// NewTaobaoOpenAccountCreateRequest 初始化TaobaoOpenAccountCreateAPIRequest对象
func NewTaobaoOpenAccountCreateRequest() *TaobaoOpenAccountCreateAPIRequest {
	return &TaobaoOpenAccountCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountCreateAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenAccountCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// Open Account的列表
func (r *TaobaoOpenAccountCreateAPIRequest) SetParamList(_paramList []OpenAccount) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TaobaoOpenAccountCreateAPIRequest) GetParamList() []OpenAccount {
	return r._paramList
}
