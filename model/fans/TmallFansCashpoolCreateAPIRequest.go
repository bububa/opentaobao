package fans

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallfanscashpoolcreateAPIRequest 创建资金池 API请求
// tmall.fans.cashpool.create
//
// 商家创建资金池接口
type TmallfanscashpoolcreateAPIRequest struct {
	model.Params
	// 创建资奖池输入对象
	_createCashPoolParamDo *CreateCashPoolParamDo
}

// NewTmallfanscashpoolcreateRequest 初始化TmallfanscashpoolcreateAPIRequest对象
func NewTmallfanscashpoolcreateRequest() *TmallfanscashpoolcreateAPIRequest {
	return &TmallfanscashpoolcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallfanscashpoolcreateAPIRequest) GetApiMethodName() string {
	return "tmall.fans.cashpool.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallfanscashpoolcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallfanscashpoolcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateCashPoolParamDo is CreateCashPoolParamDo Setter
// 创建资奖池输入对象
func (r *TmallfanscashpoolcreateAPIRequest) SetCreateCashPoolParamDo(_createCashPoolParamDo *CreateCashPoolParamDo) error {
	r._createCashPoolParamDo = _createCashPoolParamDo
	r.Set("create_cash_pool_param_do", _createCashPoolParamDo)
	return nil
}

// GetCreateCashPoolParamDo CreateCashPoolParamDo Getter
func (r TmallfanscashpoolcreateAPIRequest) GetCreateCashPoolParamDo() *CreateCashPoolParamDo {
	return r._createCashPoolParamDo
}
