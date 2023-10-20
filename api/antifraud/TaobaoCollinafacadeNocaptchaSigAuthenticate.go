package antifraud

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/antifraud"
)

// Taobaocollinafacadenocaptchasigauthenticate 人机识别
// taobao.collinafacade.nocaptcha.sig.authenticate
//
// 人机识别颁发签名串后,本接口负责向ISV提供签名串校验服务
func Taobaocollinafacadenocaptchasigauthenticate(clt *core.SDKClient, req *antifraud.TaobaocollinafacadenocaptchasigauthenticateAPIRequest, session string) (*antifraud.TaobaocollinafacadenocaptchasigauthenticateAPIResponse, error) {
	var resp antifraud.TaobaocollinafacadenocaptchasigauthenticateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
