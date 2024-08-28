package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaLoginAuthsignGet 获取登陆权限签名
// taobao.simba.login.authsign.get
//
// 获取登陆权限签名
func TaobaoSimbaLoginAuthsignGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaLoginAuthsignGetAPIRequest, resp *simba.TaobaoSimbaLoginAuthsignGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
