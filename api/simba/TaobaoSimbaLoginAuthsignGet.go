package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaLoginAuthsignGet 获取登陆权限签名
// taobao.simba.login.authsign.get
//
// 获取登陆权限签名
func TaobaoSimbaLoginAuthsignGet(clt *core.SDKClient, req *simba.TaobaoSimbaLoginAuthsignGetAPIRequest, resp *simba.TaobaoSimbaLoginAuthsignGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
