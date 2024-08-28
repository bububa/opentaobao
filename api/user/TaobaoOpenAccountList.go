package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountList OpenAccount账号信息查询
// taobao.open.account.list
//
// OpenAccount账号信息查询
func TaobaoOpenAccountList(ctx context.Context, clt *core.SDKClient, req *user.TaobaoOpenAccountListAPIRequest, resp *user.TaobaoOpenAccountListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
