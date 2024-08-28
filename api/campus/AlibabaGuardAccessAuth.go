package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaGuardAccessAuth 鉴权
// alibaba.guard.access.auth
//
// 刷卡鉴权
func AlibabaGuardAccessAuth(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaGuardAccessAuthAPIRequest, resp *campus.AlibabaGuardAccessAuthAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
