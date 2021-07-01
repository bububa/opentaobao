package antifraud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCollinafacadeNocaptchaSigAuthenticateAPIRequest
人机识别 API请求
taobao.collinafacade.nocaptcha.sig.authenticate

人机识别颁发签名串后,本接口负责向ISV提供签名串校验服务 */
type TaobaoCollinafacadeNocaptchaSigAuthenticateAPIRequest struct {
	model.Params
	// 签名串校验接口入参
	_sigAuthenticateContext *SigAuthenticateContext
}

// NewTaobaoCollinafacadeNocaptchaSigAuthenticateRequest 初始化TaobaoCollinafacadeNocaptchaSigAuthenticateAPIRequest对象
func NewTaobaoCollinafacadeNocaptchaSigAuthenticateRequest() *TaobaoCollinafacadeNocaptchaSigAuthenticateAPIRequest {
	return &TaobaoCollinafacadeNocaptchaSigAuthenticateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCollinafacadeNocaptchaSigAuthenticateAPIRequest) GetApiMethodName() string {
	return "taobao.collinafacade.nocaptcha.sig.authenticate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCollinafacadeNocaptchaSigAuthenticateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SigAuthenticateContext Setter
// 签名串校验接口入参
func (r *TaobaoCollinafacadeNocaptchaSigAuthenticateAPIRequest) SetSigAuthenticateContext(_sigAuthenticateContext *SigAuthenticateContext) error {
	r._sigAuthenticateContext = _sigAuthenticateContext
	r.Set("sig_authenticate_context", _sigAuthenticateContext)
	return nil
}

// Get SigAuthenticateContext Getter
func (r TaobaoCollinafacadeNocaptchaSigAuthenticateAPIRequest) GetSigAuthenticateContext() *SigAuthenticateContext {
	return r._sigAuthenticateContext
}
