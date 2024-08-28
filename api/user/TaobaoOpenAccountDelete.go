package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountDelete OpenAccount删除数据
// taobao.open.account.delete
//
// OpenAccount删除数据
func TaobaoOpenAccountDelete(ctx context.Context, clt *core.SDKClient, req *user.TaobaoOpenAccountDeleteAPIRequest, resp *user.TaobaoOpenAccountDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
