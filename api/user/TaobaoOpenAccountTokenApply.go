package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountTokenApply 申请免登Open Account Token
// taobao.open.account.token.apply
//
// 申请免登Open Account Token
func TaobaoOpenAccountTokenApply(ctx context.Context, clt *core.SDKClient, req *user.TaobaoOpenAccountTokenApplyAPIRequest, resp *user.TaobaoOpenAccountTokenApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
