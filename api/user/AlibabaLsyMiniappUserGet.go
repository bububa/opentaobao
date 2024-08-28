package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaLsyMiniappUserGet 零售云小程序获取登录用户信息
// alibaba.lsy.miniapp.user.get
//
// 零售云小程序，通过授权码获取登录的卖家账号信息
func AlibabaLsyMiniappUserGet(ctx context.Context, clt *core.SDKClient, req *user.AlibabaLsyMiniappUserGetAPIRequest, resp *user.AlibabaLsyMiniappUserGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
