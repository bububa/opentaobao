package yunosaccount

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosaccount"
)

// YunosAccountCallapi 调用YunOS账号开放API
// yunos.account.callapi
//
// YunOS账号客户端对外开放的api通过top暴露
func YunosAccountCallapi(ctx context.Context, clt *core.SDKClient, req *yunosaccount.YunosAccountCallapiAPIRequest, resp *yunosaccount.YunosAccountCallapiAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
