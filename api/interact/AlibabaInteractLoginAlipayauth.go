package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractLoginAlipayauth 双11到店互动花呗红包获取token鉴权接口
// alibaba.interact.login.alipayauth
//
// 双11到店互动花呗红包获取token鉴权接口
func AlibabaInteractLoginAlipayauth(clt *core.SDKClient, req *interact.AlibabaInteractLoginAlipayauthAPIRequest, resp *interact.AlibabaInteractLoginAlipayauthAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
