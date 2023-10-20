package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractloginalipayauth 双11到店互动花呗红包获取token鉴权接口
// alibaba.interact.login.alipayauth
//
// 双11到店互动花呗红包获取token鉴权接口
func Alibabainteractloginalipayauth(clt *core.SDKClient, req *interact.AlibabainteractloginalipayauthAPIRequest, session string) (*interact.AlibabainteractloginalipayauthAPIResponse, error) {
	var resp interact.AlibabainteractloginalipayauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
