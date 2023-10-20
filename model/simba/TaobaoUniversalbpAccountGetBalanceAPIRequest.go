package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpaccountgetbalanceAPIRequest 获取账户余额，现金余额 API请求
// taobao.universalbp.account.get.balance
//
// 获取账户实时现金余额
type TaobaouniversalbpaccountgetbalanceAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaouniversalbpaccountgetbalanceRequest 初始化TaobaouniversalbpaccountgetbalanceAPIRequest对象
func NewTaobaouniversalbpaccountgetbalanceRequest() *TaobaouniversalbpaccountgetbalanceAPIRequest {
	return &TaobaouniversalbpaccountgetbalanceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpaccountgetbalanceAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.account.get.balance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpaccountgetbalanceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpaccountgetbalanceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpaccountgetbalanceAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpaccountgetbalanceAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
