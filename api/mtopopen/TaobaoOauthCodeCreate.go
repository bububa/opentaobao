package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoOauthCodeCreate 淘宝OauthCode颁发
// taobao.oauth.code.create
//
// 手淘无线开放的oauthCode颁发接口
func TaobaoOauthCodeCreate(clt *core.SDKClient, req *mtopopen.TaobaoOauthCodeCreateAPIRequest, resp *mtopopen.TaobaoOauthCodeCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
