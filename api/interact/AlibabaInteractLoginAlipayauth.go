package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

/* AlibabaInteractLoginAlipayauth
双11到店互动花呗红包获取token鉴权接口
alibaba.interact.login.alipayauth

双11到店互动花呗红包获取token鉴权接口 */
func AlibabaInteractLoginAlipayauth(clt *core.SDKClient, req *interact.AlibabaInteractLoginAlipayauthAPIRequest, session string) (*interact.AlibabaInteractLoginAlipayauthAPIResponse, error) {
	var resp interact.AlibabaInteractLoginAlipayauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
