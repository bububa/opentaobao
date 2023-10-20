package antifraud

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/antifraud"
)

// TaobaoCollinafacadeNocaptchaSigAuthenticate 人机识别
// taobao.collinafacade.nocaptcha.sig.authenticate
//
// 人机识别颁发签名串后,本接口负责向ISV提供签名串校验服务
func TaobaoCollinafacadeNocaptchaSigAuthenticate(clt *core.SDKClient, req *antifraud.TaobaoCollinafacadeNocaptchaSigAuthenticateAPIRequest, resp *antifraud.TaobaoCollinafacadeNocaptchaSigAuthenticateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
