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

// New
